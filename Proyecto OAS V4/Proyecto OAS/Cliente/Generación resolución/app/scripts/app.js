'use strict';

/**
 * @ngdoc overview
 * @name generacionResolucionApp
 * @description
 * # generacionResolucionApp
 *
 * Main module of the application.
 */
angular
  .module('generacionResolucionApp', [
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
    'resolucion_service'
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
      .when('/resolucion_pregrado', {
        templateUrl: 'views/resolucion_pregrado.html',
        controller: 'ResolucionPregradoCtrl',
        controllerAs: 'resolucionPregrado'
      })
      .when('/resolucion_posgrado', {
        templateUrl: 'views/resolucion_posgrado.html',
        controller: 'ResolucionPosgradoCtrl',
        controllerAs: 'resolucionPosgrado'
      })
      .when('/Resolucion_pregrado_TCO_MTO', {
        templateUrl: 'views/resolucion_pregrado_tco_mto.html',
        controller: 'ResolucionPregradoTcoMtoCtrl',
        controllerAs: 'ResolucionPregradoTCOMTO'
      })
      .when('/Resolucion_pregrado_HCP', {
        templateUrl: 'views/resolucion_pregrado_hcp.html',
        controller: 'ResolucionPregradoHcpCtrl',
        controllerAs: 'ResolucionPregradoHCP'
      })
      .when('/Resolucion_pregrado_HCH', {
        templateUrl: 'views/resolucion_pregrado_hch.html',
        controller: 'ResolucionPregradoHchCtrl',
        controllerAs: 'ResolucionPregradoHCH'
      })
      .when('/Resolucion_posgrado_HCP', {
        templateUrl: 'views/resolucion_posgrado_hcp.html',
        controller: 'ResolucionPosgradoHcpCtrl',
        controllerAs: 'ResolucionPosgradoHCP'
      })
      .when('/Resolucion_posgrado_HCH', {
        templateUrl: 'views/resolucion_posgrado_hch.html',
        controller: 'ResolucionPosgradoHchCtrl',
        controllerAs: 'ResolucionPosgradoHCH'
      })
      .otherwise({
        redirectTo: '/'
      });
  }]);
