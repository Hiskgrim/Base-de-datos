// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"appAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/funcionario",
			beego.NSInclude(
				&controllers.FuncionarioController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/categoria",
			beego.NSInclude(
				&controllers.CategoriaController{},
			),
		),

		beego.NSNamespace("/facultad",
			beego.NSInclude(
				&controllers.FacultadController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/estado_contrato",
			beego.NSInclude(
				&controllers.EstadoContratoController{},
			),
		),

		beego.NSNamespace("/tipo_contratacion",
			beego.NSInclude(
				&controllers.TipoContratacionController{},
			),
		),

		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/docente",
			beego.NSInclude(
				&controllers.DocenteController{},
			),
		),

		beego.NSNamespace("/carga_academica",
			beego.NSInclude(
				&controllers.CargaAcademicaController{},
			),
		),

		beego.NSNamespace("/contrato_especial",
			beego.NSInclude(
				&controllers.ContratoEspecialController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),

		beego.NSNamespace("/cancelacion_contrato",
			beego.NSInclude(
				&controllers.CancelacionContratoController{},
			),
		),

		beego.NSNamespace("/solicitud_necesidad",
			beego.NSInclude(
				&controllers.SolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/clasificacion",
			beego.NSInclude(
				&controllers.ClasificacionController{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/servicio",
			beego.NSInclude(
				&controllers.ServicioController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/modalidad_seleccion",
			beego.NSInclude(
				&controllers.ModalidadSeleccionController{},
			),
		),

		beego.NSNamespace("/estado_solicitud",
			beego.NSInclude(
				&controllers.EstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/estado_necesidad_necesidad",
			beego.NSInclude(
				&controllers.EstadoNecesidadNecesidadController{},
			),
		),

		beego.NSNamespace("/necesidad",
			beego.NSInclude(
				&controllers.NecesidadController{},
			),
		),

		beego.NSNamespace("/asignatura",
			beego.NSInclude(
				&controllers.AsignaturaController{},
			),
		),

		beego.NSNamespace("/grupo",
			beego.NSInclude(
				&controllers.GrupoController{},
			),
		),

		beego.NSNamespace("/docente_grupo",
			beego.NSInclude(
				&controllers.DocenteGrupoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
