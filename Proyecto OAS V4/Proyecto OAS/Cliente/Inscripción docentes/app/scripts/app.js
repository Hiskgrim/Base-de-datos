'use strict';

/**
 * @ngdoc overview
 * @name inscripcionDocentesApp
 * @description
 * # inscripcionDocentesApp
 *
 * Main module of the application.
 */
angular
  .module('inscripcionDocentesApp', [
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
    'inscripcion_service'
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
      .when('/inscripcion_pregrado', {
        templateUrl: 'views/inscripcion_pregrado.html',
        controller: 'InscripcionPregradoCtrl',
        controllerAs: 'inscripcionPregrado'
      })
      .when('/inscripcion_posgrado', {
        templateUrl: 'views/inscripcion_posgrado.html',
        controller: 'InscripcionPosgradoCtrl',
        controllerAs: 'inscripcionPosgrado'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
