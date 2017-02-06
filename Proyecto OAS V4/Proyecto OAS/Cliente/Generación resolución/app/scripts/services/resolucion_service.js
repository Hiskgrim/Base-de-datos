'use strict';

/**
 * @ngdoc service
 * @name generacionResolucionApp.resolucionService
 * @description
 * # resolucionService
 * Factory in the generacionResolucionApp.
 */
angular.module('resolucion_service',[])
  .factory('resolucion_request', function ($http) {
    // Service logic
    // ...
    var path = "http://localhost:8080/v1/";

    // Public API here
    return {
      getAll: function (table) {
        return $http.get(path+table);
      },
      post: function (table,elemento) {
        return $http.post(path+table,elemento);
      },
      delete: function (table,id) {
        return $http.delete(path+table+"/"+id);
      },
      getOne: function (table,id) {
        return $http.get(path+table+"/"+id);
      },
      put: function (table,id,elemento) {
        return $http.put(path+table+"/"+id,elemento);
      }
    };
  });
