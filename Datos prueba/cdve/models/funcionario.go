package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Funcionario struct {
	Id              int       `orm:"column(id);pk;auto"`
	Documento       string    `orm:"column(documento)"`
	IdCiudad        *Ciudad   `orm:"column(id_ciudad);rel(fk)"`
	Direccion       string    `orm:"column(direccion)"`
	TelefonoFijo    string    `orm:"column(telefono_fijo);null"`
	TelefonoCelular string    `orm:"column(telefono_celular);null"`
	Email           string    `orm:"column(email)"`
	PrimerNombre    string    `orm:"column(primer_nombre)"`
	SegundoNombre   string    `orm:"column(segundo_nombre);null"`
	PrimerApellido  string    `orm:"column(primer_apellido)"`
	SegundoApellido string    `orm:"column(segundo_apellido);null"`
	Sexo            string    `orm:"column(sexo)"`
	FechaNacimiento time.Time `orm:"column(fecha_nacimiento);type(date)"`
	Activo          bool      `orm:"column(activo)"`
	Rut             int       `orm:"column(rut)"`
}

func (t *Funcionario) TableName() string {
	return "funcionario"
}

func init() {
	orm.RegisterModel(new(Funcionario))
}

// AddFuncionario insert a new Funcionario into database and returns
// last inserted Id on success.
func AddFuncionario(m *Funcionario) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFuncionarioById retrieves Funcionario by Id. Returns error if
// Id doesn't exist
func GetFuncionarioById(id int) (v *Funcionario, err error) {
	o := orm.NewOrm()
	v = &Funcionario{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFuncionario retrieves all Funcionario matches certain condition. Returns empty list if
// no records exist
func GetAllFuncionario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Funcionario))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Funcionario
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateFuncionario updates Funcionario by Id and returns error if
// the record to be updated doesn't exist
func UpdateFuncionarioById(m *Funcionario) (err error) {
	o := orm.NewOrm()
	v := Funcionario{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFuncionario deletes Funcionario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFuncionario(id int) (err error) {
	o := orm.NewOrm()
	v := Funcionario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Funcionario{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
