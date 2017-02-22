'use strict';

/**
 * @ngdoc overview
 * @name cdveApp
 * @description
 * # cdveApp
 *
 * Main module of the application.
 */
angular
  .module('cdveApp', [
    'ngAnimate',
    'ngResource',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'afOAuth2',
    'treeControl',
    'ngMaterial',
    'ui.grid',
    'ui.grid.edit',
    'ui.grid.rowEdit',
    'ui.grid.cellNav',
    'ui.grid.treeView',
    'ui.grid.selection',
    'ui.grid.exporter',
    'ngStorage',
    'vinculacion_service'
  ])
    .config(['$locationProvider','$routeProvider', function($locationProvider, $routeProvider) {
      $locationProvider.hashPrefix("");
      $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl',
        controllerAs: 'main'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl',
        controllerAs: 'about'
      })
      .when('/inscripcionPregrado', {
        templateUrl: 'views/inscripcionpregrado.html',
        controller: 'InscripcionpregradoCtrl',
        controllerAs: 'inscripcionPregrado'
      })
      .when('/inscripcionPosgrado', {
        templateUrl: 'views/inscripcionposgrado.html',
        controller: 'InscripcionposgradoCtrl',
        controllerAs: 'inscripcionPosgrado'
      })
      .when('/cargaPregrado', {
        templateUrl: 'views/cargapregrado.html',
        controller: 'CargapregradoCtrl',
        controllerAs: 'cargaPregrado'
      })
      .when('/cargaPosgrado', {
        templateUrl: 'views/cargaposgrado.html',
        controller: 'CargaposgradoCtrl',
        controllerAs: 'cargaPosgrado'
      })
      .when('/cargaDocente/:IdDocente', {
        templateUrl: 'views/cargadocente.html',
        controller: 'CargadocenteCtrl',
        controllerAs: 'cargaDocente'
      })
      .when('/resolucionpregradoHCH', {
        templateUrl: 'views/resolucionpregradohch.html',
        controller: 'ResolucionpregradohchCtrl',
        controllerAs: 'resolucionPregradoHCH'
      })
      .when('/resolucionpregradoHCP', {
        templateUrl: 'views/resolucionpregradohcp.html',
        controller: 'ResolucionpregradohcpCtrl',
        controllerAs: 'resolucionPregradoHCP'
      })
      .when('/resolucionpregradoTCOMTO', {
        templateUrl: 'views/resolucionpregradotcomto.html',
        controller: 'ResolucionpregradotcomtoCtrl',
        controllerAs: 'resolucionPregradoTCOMTO'
      })
      .when('/resolucionposgradoHCH', {
        templateUrl: 'views/resolucionposgradohch.html',
        controller: 'ResolucionposgradohchCtrl',
        controllerAs: 'resolucionPosgradoHCH'
      })
      .when('/resolucionposgradoHCP', {
        templateUrl: 'views/resolucionposgradohcp.html',
        controller: 'ResolucionposgradohcpCtrl',
        controllerAs: 'resolucionPosgradoHCP'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
