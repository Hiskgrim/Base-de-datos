'use strict';

/**
 * @ngdoc service
 * @name inscripcionDocentesApp.inscripcionService
 * @description
 * # inscripcionService
 * Factory in the inscripcionDocentesApp.
 */
angular.module('inscripcion_service',[])
  .factory('inscripcion_request', function ($http) {
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
