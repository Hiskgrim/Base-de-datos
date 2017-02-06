'use strict';

/**
 * @ngdoc function
 * @name inscripcionDocentesApp.controller:InscripcionPregradoCtrl
 * @description
 * # InscripcionPregradoCtrl
 * Controller of the inscripcionDocentesApp
 */
angular.module('inscripcionDocentesApp')
  .controller('InscripcionPregradoCtrl', function (inscripcion_request) {
    var self=this;

    inscripcion_request.getAll("facultad").then(function(response){
      self.facultades=response.data;
    });

    inscripcion_request.getAll("proyecto_curricular").then(function(response){
      self.proyectos=response.data;
    });

    inscripcion_request.getAll("categoria").then(function(response){
      self.categorias=response.data;
    });

    inscripcion_request.getAll("dedicacion").then(function(response){
      self.dedicaciones=response.data;
    });

    inscripcion_request.getAll("ciudad").then(function(response){
      self.ciudades=response.data;
    });

    inscripcion_request.getAll("docente").then(function(response){
      self.docentes=response.data;
    });

    inscripcion_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    inscripcion_request.getAll("clasificacion").then(function(response){
      self.clasificaciones=response.data;
    });

    inscripcion_request.getAll("contrato_general").then(function(response){
      self.contratos=response.data;
    });

    inscripcion_request.getAll("contrato_especial").then(function(response){
      self.contratosEspeciales=response.data;
    });

    inscripcion_request.getAll("carga_academica").then(function(response){
      self.cargasAcademicas=response.data;
    });

    self.agregarInformacion = function(datos){      
      self.agregarFuncionario(datos);
    };

    self.agregarFuncionario = function(datos){
      var funcionariodatos = {
        PrimerNombre: datos.primernombre,
        SegundoNombre: datos.segundonombre,
        PrimerApellido: datos.primerapellido,
        SegundoApellido: datos.segundoapellido,
        PrimerNombre: datos.primernombre,
        Documento: datos.documento,
        IdCiudad: {
          Id: parseInt(datos.ciudadexpedicion)
        },
        Sexo: datos.sexo,
        FechaNacimiento: datos.fechanacimiento,
        Rut: datos.rut,
        Telefono: datos.telefono,
        Direccion: datos.direccion,
        Email: datos.email,
        Activo: false
      };
      inscripcion_request.post("funcionario",funcionariodatos).then(function(response){
        datos.funcionario=response.data.Id;
        self.agregarDocente(datos);
        self.agregarContrato(datos);
      });
    };

    self.agregarDocente = function(datos){
      var docentedatos = {
        IdFuncionario: {
          Id: datos.funcionario
        }
      };
      inscripcion_request.post("docente",docentedatos).then(function(response){
        datos.docente=response.data.Id;
        self.agregarClasificacion(datos);
        self.agregarCargaAcademica(datos);
      });
    };

    self.agregarClasificacion = function(datos){
      var clasificaciondatos = {
        FechaRegistro: new Date(),
        IdDedicacion: {
          Id: parseInt(datos.dedicacion)
        },
        IdCategoria: {
          Id: parseInt(datos.categoria)
        },
        IdDocente: {
          Id: datos.docente
        }
      };
      inscripcion_request.post("clasificacion",clasificaciondatos);
    };

    self.agregarCargaAcademica = function(datos){
      var cargaAcademicadatos = {
        HorasSemanales: datos.horassemanales,
        IdProyectoCurricular: {
          Id: parseInt(datos.proyectocurricular)
        },
        IdDocente: {
          Id: datos.docente
        }
      };
      inscripcion_request.post("carga_academica",cargaAcademicadatos);
    };

    self.agregarContrato = function(datos){
      var contratodatos = {
        Vigencia: datos.semanas,
        IdTipoContrato: {
          Id: 1
        }
      };
      self.agregarContratoEspecial(datos);
      inscripcion_request.post("contrato_general",contratodatos).then(function(response){
        datos.contrato=response.data.Id;
        self.agregarContratoEspecial(datos);
      });
    };

    self.agregarContratoEspecial = function(datos){
      var contratoEspecialdatos = {
        Vigencia: datos.semanas,
        NumeroContratoGeneral: {
          Id: datos.contrato
        },
        IdFuncionario: {
          Id: datos.funcionario
        },
        NumeroResolucion: 1,
        AnioResolucion: 2017
      };
      inscripcion_request.post("contrato_especial",contratoEspecialdatos);
    };

  });
