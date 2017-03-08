'use strict';

/**
 * @ngdoc function
 * @name cdveApp.controller:CargapregradoCtrl
 * @description
 * # CargapregradoCtrl
 * Controller of the cdveApp
 */
angular.module('cdveApp')
  .controller('CargapregradoCtrl', function (vinculacion_request, $location) {
    var self = this;

    vinculacion_request.getAll("proyecto_curricular").then(function(response){
      self.proyectos=response.data;
    });

    vinculacion_request.getAll("categoria").then(function(response){
      self.categorias=response.data;
    });

    vinculacion_request.getAll("dedicacion").then(function(response){
      self.dedicaciones=response.data;
    });

    vinculacion_request.getAll("ciudad").then(function(response){
      self.ciudades=response.data;
    });

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
      self.contratosEspeciales=response.data;
    });

    vinculacion_request.getAll("carga_academica").then(function(response){
      self.cargasAcademicas=response.data;
    });

    self.abrirUbicacion=function(idDocente){
    	$location.url('/cargaDocente/' + idDocente);
    }
  });
