package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:AsignaturaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CancelacionContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CargaAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CategoriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CiudadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CiudadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CiudadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CiudadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:CiudadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:DocenteGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoNecesidadNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["appAPI/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FacultadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FacultadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FacultadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FacultadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FacultadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:FuncionarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:GrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:GrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:GrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:GrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:GrupoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ServicioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ServicioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ServicioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ServicioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:ServicioController"] = append(beego.GlobalControllerRouter["appAPI/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["appAPI/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["appAPI/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["appAPI/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
