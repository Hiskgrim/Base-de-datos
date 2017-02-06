'use strict';

/**
 * @ngdoc overview
 * @name asignacionCargaApp
 * @description
 * # asignacionCargaApp
 *
 * Main module of the application.
 */
angular
  .module('asignacionCargaApp', [
    'ngAnimate',
    'ngCookies',
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
    'carga_service'
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
      .when('/carga_pregrado', {
        templateUrl: 'views/carga_pregrado.html',
        controller: 'CargaPregradoCtrl',
        controllerAs: 'cargaPregrado'
      })
      .when('/carga_posgrado', {
        templateUrl: 'views/carga_posgrado.html',
        controller: 'CargaPosgradoCtrl',
        controllerAs: 'cargaPosgrado'
      })
      .when('/carga_docente/:IdDocente', {
        templateUrl: 'views/carga_docente.html',
        controller: 'CargaDocenteCtrl',
        controllerAs: 'cargaDocente'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
