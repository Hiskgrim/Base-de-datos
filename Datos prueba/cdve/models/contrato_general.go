package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ContratoGeneral struct {
	Id                           int           `orm:"column(numero);pk;auto"`
	Vigencia                     int           `orm:"column(vigencia);null"`
	ObjetoContrato               string        `orm:"column(objeto_contrato);null"`
	PlazoEjecución               int           `orm:"column(plazo_ejecución);null"`
	ClausulaRegistroPresupuestal bool          `orm:"column(clausula_registro_presupuestal);null"`
	SedeSolicitante              string        `orm:"column(sede_solicitante);null"`
	DependenciaSolicitante       string        `orm:"column(dependencia_solicitante);null"`
	NumeroSolicitudNecesidad     int           `orm:"column(numero_solicitud_necesidad);null"`
	NumeroCdp                    int16         `orm:"column(numero_cdp);null"`
	Contratista                  float64       `orm:"column(contratista);null"`
	ValorContrato                float64       `orm:"column(valor_contrato);null"`
	Justificacion                string        `orm:"column(justificacion);null"`
	DescripcionFormaPago         string        `orm:"column(descripcion_forma_pago);null"`
	Condiciones                  string        `orm:"column(condiciones);null"`
	FechaRegistro                time.Time     `orm:"column(fecha_registro);type(date);null"`
	TipologiaContrato            int           `orm:"column(tipologia_contrato);null"`
	TipoCompromiso               int           `orm:"column(tipo_compromiso);null"`
	ModalidadSeleccion           int           `orm:"column(modalidad_seleccion);null"`
	Procedimiento                int           `orm:"column(procedimiento);null"`
	RegimenContratacion          int           `orm:"column(regimen_contratacion);null"`
	TipoGasto                    int           `orm:"column(tipo_gasto);null"`
	TomaGastoInversion           int           `orm:"column(toma_gasto_inversion);null"`
	OrigenPresupuesto            int           `orm:"column(origen_presupuesto);null"`
	OrigenRecursos               int           `orm:"column(origen_recursos);null"`
	TipoMoneda                   int           `orm:"column(tipo_moneda);null"`
	ValorContratoMe              float64       `orm:"column(valor_contrato_me);null"`
	ValorTasaCambio              float64       `orm:"column(valor_tasa_cambio);null"`
	TipoControl                  int           `orm:"column(tipo_control);null"`
	Observaciones                string        `orm:"column(observaciones);null"`
	ClaseContratista             int           `orm:"column(clase_contratista);null"`
	Convenio                     string        `orm:"column(convenio);null"`
	NumeroConstancia             int           `orm:"column(numero_constancia);null"`
	Estado                       bool          `orm:"column(estado);null"`
	RegistroPresupuestal         int           `orm:"column(registro_presupuestal);null"`
	IdTipoContrato               *TipoContrato `orm:"column(id_tipo_contrato);rel(fk)"`
}

func (t *ContratoGeneral) TableName() string {
	return "contrato_general"
}

func init() {
	orm.RegisterModel(new(ContratoGeneral))
}

// AddContratoGeneral insert a new ContratoGeneral into database and returns
// last inserted Id on success.
func AddContratoGeneral(m *ContratoGeneral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratoGeneralById retrieves ContratoGeneral by Id. Returns error if
// Id doesn't exist
func GetContratoGeneralById(id int) (v *ContratoGeneral, err error) {
	o := orm.NewOrm()
	v = &ContratoGeneral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContratoGeneral retrieves all ContratoGeneral matches certain condition. Returns empty list if
// no records exist
func GetAllContratoGeneral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContratoGeneral))
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

	var l []ContratoGeneral
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

// UpdateContratoGeneral updates ContratoGeneral by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratoGeneralById(m *ContratoGeneral) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContratoGeneral deletes ContratoGeneral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContratoGeneral(id int) (err error) {
	o := orm.NewOrm()
	v := ContratoGeneral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContratoGeneral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
