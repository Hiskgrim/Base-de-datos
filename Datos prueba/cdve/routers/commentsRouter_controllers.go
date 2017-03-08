package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["cdve/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:CiudadController"] = append(beego.GlobalControllerRouter["cdve/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DocenteController"] = append(beego.GlobalControllerRouter["cdve/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DocenteController"] = append(beego.GlobalControllerRouter["cdve/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DocenteController"] = append(beego.GlobalControllerRouter["cdve/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DocenteController"] = append(beego.GlobalControllerRouter["cdve/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:DocenteController"] = append(beego.GlobalControllerRouter["cdve/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FacultadController"] = append(beego.GlobalControllerRouter["cdve/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["cdve/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"] = append(beego.GlobalControllerRouter["cdve/controllers:TipoProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
