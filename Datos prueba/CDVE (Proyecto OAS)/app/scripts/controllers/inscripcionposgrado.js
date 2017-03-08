'use strict';

/**
 * @ngdoc function
 * @name cdveApp.controller:InscripcionposgradoCtrl
 * @description
 * # InscripcionposgradoCtrl
 * Controller of the cdveApp
 */
angular.module('cdveApp')
  .controller('InscripcionposgradoCtrl', function (vinculacion_request,$routeParams) {
    var self=this;

    self.IdFacultad=parseInt($routeParams.IdFacultad);

    vinculacion_request.getAll("proyecto_curricular","query=IdFacultad.id%3A"+self.IdFacultad+"%2CIdTipoProyectoCurricular.Id%3A2&limit=0").then(function(response){
      self.proyectos=response.data;
    });

    vinculacion_request.getAll("categoria","limit=0").then(function(response){
      self.categorias=response.data;
    });

    vinculacion_request.getAll("dedicacion","limit=0").then(function(response){
      self.dedicaciones=response.data;
    });

    vinculacion_request.getAll("ciudad","limit=0").then(function(response){
      self.ciudades=response.data;
    });
/*
    vinculacion_request.getAll("docente").then(function(response){
      self.docentes=response.data;
    });

    vinculacion_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    vinculacion_request.getAll("clasificacion").then(function(response){
      self.clasificaciones=response.data;
    });

    vinculacion_request.getAll("contrato_general").then(function(response){
      self.contratos=response.data;
    });

    vinculacion_request.getAll("contrato_especial").then(function(response){
      self.contratoEspecial=response.data;
    });

    vinculacion_request.getAll("carga_academica").then(function(response){
      self.cargasAcademicas=response.data;
    });
*/


    vinculacion_request.getAll("contrato_especial","query=IdFacultad.id%3A"+self.IdFacultad+"&limit=0").then(function(response){
      var contratoEspecial=response.data;
      self.datosContratos=[];
      self.proyectos.forEach(function(proyecto){
        self.datosContratos[proyecto.Id]=[];
      });
      if(contratoEspecial!=null){
        contratoEspecial.forEach(function(contratoEspecial){
          self.getContratoEspecial(contratoEspecial);
        });
      }
    });

    self.getContratoEspecial = function(contratoEspecial){
      var datosContrato={
        Nombre: null,
        Cedula: null,
        Expedida: null,
        Categoria: null,
        Dedicacion: null,
        HorasSemanales: null,
        PeriodoVinculacion: null,
        ValorContrato: null,
        ProyectoCurricular: null
      }
      vinculacion_request.getOne("contrato_general",contratoEspecial.NumeroContratoGeneral.Id).then(function(response){
        var contratoGeneral=response.data;
        datosContrato.ValorContrato=contratoGeneral.ValorContrato;
        datosContrato.PeriodoVinculacion=contratoGeneral.Vigencia;
      });
      vinculacion_request.getOne("funcionario",contratoEspecial.IdFuncionario.Id).then(function(response){
          var funcionario=response.data;
          datosContrato.Nombre=funcionario.PrimerNombre+" "+funcionario.SegundoNombre+" "+funcionario.PrimerApellido+" "+funcionario.SegundoApellido;
          datosContrato.Cedula=funcionario.Documento;
          vinculacion_request.getOne("ciudad",funcionario.IdCiudad.Id).then(function(response){
            var ciudad=response.data;
            datosContrato.Expedida=ciudad.Nombre;
          });
          vinculacion_request.getAll("docente","query=IdFuncionario.Id%3A"+funcionario.Id.toString()).then(function(response){
            var docente=response.data[0];
            vinculacion_request.getAll("clasificacion","query=IdDocente.Id%3A"+docente.Id.toString()).then(function(response){
              var clasificacion=response.data[0];
              vinculacion_request.getOne("categoria",clasificacion.IdCategoria.Id).then(function(response){
                var categoria=response.data;
                datosContrato.Categoria=categoria.Nombre;
              });
              vinculacion_request.getOne("dedicacion",clasificacion.IdDedicacion.Id).then(function(response){
                var dedicacion=response.data;
                datosContrato.Dedicacion=dedicacion.Nombre;
              });
              vinculacion_request.getAll("carga_academica","query=IdDocente.Id%3A"+docente.Id.toString()).then(function(response){
                var cargaAcademica=response.data[0];
                datosContrato.HorasSemanales=cargaAcademica.HorasSemanales;
                self.datosContratos[cargaAcademica.IdProyectoCurricular.Id].push(datosContrato);
              });
            });
          });
        });
    };

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
        Documento: datos.documento.toString(),
        IdCiudad: {
          Id: parseInt(datos.ciudadexpedicion)
        },
        Sexo: datos.sexo,
        FechaNacimiento: new Date(),
        Rut: datos.rut,
        Telefono: datos.telefono,
        Direccion: datos.direccion,
        Email: datos.email,
        Activo: false
      };
      vinculacion_request.post("funcionario",funcionariodatos).then(function(response){
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
      vinculacion_request.post("docente",docentedatos).then(function(response){
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
      vinculacion_request.post("clasificacion",clasificaciondatos);
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
      vinculacion_request.post("carga_academica",cargaAcademicadatos);
    };

    self.agregarContrato = function(datos){
      var contratodatos = {
        Vigencia: datos.semanas,
        IdTipoContrato: {
          Id: 1
        },
        ValorContrato: datos.valorcontrato
      };
      vinculacion_request.post("contrato_general",contratodatos).then(function(response){
        datos.contrato=response.data.Id;
        self.agregarContratoEspecial(datos);
      });
    };

    self.agregarContratoEspecial = function(datos){
      
      var contratoEspecialdatos = {
        Vigencia: datos.semanas,
        FechaInicio: new Date(),
        NumeroContratoGeneral: {
          Id: datos.contrato
        },
        IdFuncionario: {
          Id: datos.funcionario
        },
        IdFacultad: {
          Id: datos.facultad
        }
      };
      vinculacion_request.post("contrato_especial",contratoEspecialdatos).then(function(response){
        alert(JSON.stringify(response.data));
      });
    };
  });
