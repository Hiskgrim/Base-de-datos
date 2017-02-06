package models

type Resolucion struct {
	Numero      int       `orm:"column(numero)"`
	Anio        int16     `orm:"column(anio)"`
	Descripcion string    `orm:"column(descripcion);null"`
	Entidad     string    `orm:"column(entidad);null"`
	IdFacultad  *Facultad `orm:"column(id_facultad);rel(fk)"`
}
