'use strict';

/**
 * @ngdoc function
 * @name solicitudConvocatoriaApp.controller:SolicitudPregradoCtrl
 * @description
 * # SolicitudPregradoCtrl
 * Controller of the solicitudConvocatoriaApp
 */
angular.module('solicitudConvocatoriaApp')
  .controller('SolicitudPregradoCtrl', function (solicitud_request) {
	var self=this;

    solicitud_request.getAll("dependencia").then(function(response){
      self.dependencias=response.data;
    });

    solicitud_request.getAll("servicio").then(function(response){
      self.servicios=response.data;
    });

    solicitud_request.getAll("dependencia").then(function(response){
      self.dependencias=response.data;
    });

    solicitud_request.getAll("modalidad_seleccion").then(function(response){
      self.modalidades=response.data;
    });

    solicitud_request.getAll("unidad_ejecutora").then(function(response){
      self.unidades=response.data;
    });

    solicitud_request.getAll("necesidad").then(function(response){
      self.necesidades=response.data;
    });

    self.agregarInformacion = function(datos){      
      self.agregarNecesidad(datos);
    };

    self.agregarNecesidad = function(datos){
      var necesidaddatos = {
      	Numero: datos.numero,
      	Vigencia: datos.vigencia,
      	Objeto: datos.objeto,
        Valor: datos.valor,
        FechaSolicitud: new Date(),
        Justificacion: datos.justificacion,
        DiasDuracion: datos.vigencia,
        UnicoPago: datos.unicopago,
        AgotarPresupuesto: datos.agotarpresupuesto,
        IdUnidadEjecutora: {
        	Id: parseInt(datos.unidad)
        },
        IdModalidadSeleccion: {
        	Id: parseInt(datos.modalidad)
        },
        IdDependencia: {
        	Id: parseInt(datos.dependenciadestino)
        },
        IdServicio: {
        	Id: parseInt(datos.servicio)
        }
      };

      alert(JSON.stringify(necesidaddatos));
      solicitud_request.post("necesidad",necesidaddatos).then(function(response){
        solicitud_request.getAll("necesidad").then(function(response){
      		self.necesidades=response.data;
    	});
      });
    };


  });
