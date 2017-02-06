'use strict';

/**
 * @ngdoc overview
 * @name solicitudConvocatoriaApp
 * @description
 * # solicitudConvocatoriaApp
 *
 * Main module of the application.
 */
angular
  .module('solicitudConvocatoriaApp', [
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
    'solicitud_service'
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
      .when('/solicitud_pregrado', {
        templateUrl: 'views/solicitud_pregrado.html',
        controller: 'SolicitudPregradoCtrl',
        controllerAs: 'solicitudPregrado'
      })
      .when('/solicitud_posgrado', {
        templateUrl: 'views/solicitud_posgrado.html',
        controller: 'SolicitudPosgradoCtrl',
        controllerAs: 'solicitudPosgrado'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
