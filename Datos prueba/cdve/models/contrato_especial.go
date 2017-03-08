package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ContratoEspecial struct {
	Id                    int              `orm:"column(id);pk;auto"`
	FechaInicio           time.Time        `orm:"column(fecha_inicio);type(date);null"`
	Vigencia              int16            `orm:"column(vigencia);null"`
	NumeroContratoGeneral *ContratoGeneral `orm:"column(numero_contrato_general);rel(fk)"`
	IdFuncionario         *Funcionario     `orm:"column(id_funcionario);rel(fk)"`
	IdFacultad            *Facultad        `orm:"column(id_facultad);rel(fk)"`
}

func (t *ContratoEspecial) TableName() string {
	return "contrato_especial"
}

func init() {
	orm.RegisterModel(new(ContratoEspecial))
}

// AddContratoEspecial insert a new ContratoEspecial into database and returns
// last inserted Id on success.
func AddContratoEspecial(m *ContratoEspecial) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratoEspecialById retrieves ContratoEspecial by Id. Returns error if
// Id doesn't exist
func GetContratoEspecialById(id int) (v *ContratoEspecial, err error) {
	o := orm.NewOrm()
	v = &ContratoEspecial{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContratoEspecial retrieves all ContratoEspecial matches certain condition. Returns empty list if
// no records exist
func GetAllContratoEspecial(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContratoEspecial))
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

	var l []ContratoEspecial
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

// UpdateContratoEspecial updates ContratoEspecial by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratoEspecialById(m *ContratoEspecial) (err error) {
	o := orm.NewOrm()
	v := ContratoEspecial{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContratoEspecial deletes ContratoEspecial by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContratoEspecial(id int) (err error) {
	o := orm.NewOrm()
	v := ContratoEspecial{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContratoEspecial{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
