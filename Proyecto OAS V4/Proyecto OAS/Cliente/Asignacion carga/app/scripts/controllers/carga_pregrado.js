'use strict';

/**
 * @ngdoc function
 * @name asignacionCargaApp.controller:CargaPregradoCtrl
 * @description
 * # CargaPregradoCtrl
 * Controller of the asignacionCargaApp
 */
angular.module('asignacionCargaApp')
  .controller('CargaPregradoCtrl', function (carga_request, $location) {
    var self = this;

    carga_request.getAll("proyecto_curricular").then(function(response){
      self.proyectos=response.data;
    });

    carga_request.getAll("categoria").then(function(response){
      self.categorias=response.data;
    });

    carga_request.getAll("dedicacion").then(function(response){
      self.dedicaciones=response.data;
    });

    carga_request.getAll("ciudad").then(function(response){
      self.ciudades=response.data;
    });

    carga_request.getAll("docente").then(function(response){
      self.docentes=response.data;
    });

    carga_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    carga_request.getAll("clasificacion").then(function(response){
      self.clasificaciones=response.data;
    });

    carga_request.getAll("contrato_general").then(function(response){
      self.contratos=response.data;
    });

    carga_request.getAll("contrato_especial").then(function(response){
      self.contratosEspeciales=response.data;
    });

    carga_request.getAll("carga_academica").then(function(response){
      self.cargasAcademicas=response.data;
    });

    self.abrirUbicacion=function(idDocente){
    	$location.url('/carga_docente/' + idDocente);
    }

  });
