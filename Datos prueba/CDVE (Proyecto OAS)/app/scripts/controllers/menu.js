'use strict';

angular.module('cdveApp')
  .controller('menuCtrl', function($location, $http, $scope, token_service, vinculacion_request) {
    var ctrl = this;
    vinculacion_request.getAll("facultad","limit=0").then(function(response){
      $scope.facultades=response.data;
      var contenidoInscripcion=[];
      var contenidoResolucion=[];
      $scope.facultades.forEach(function(facultad){
        var itemInscripcion={
          "Id": 1,
          "Nombre": facultad.Nombre,
          "Opciones":[{
            "Id": 2,
            "Nombre": "Pregrado",
            "Url": "inscripcionPregrado/"+facultad.Id.toString(),
            "Opciones": null
          },{
            "Id": 4,
            "Nombre": "Posgrado",
            "Url": "inscripcionPosgrado/"+facultad.Id.toString(),
            "Opciones": null
          }]          
        }
        var itemResolucion={
          "Id": 1,
          "Nombre": facultad.Nombre,
          "Opciones":[{
            "Id": 7,
            "Nombre": "Pregrado",
            "Url": "urnl_nivel_2",
            "Opciones": [{
              "Id": 16,
              "Nombre": "TCO - MTO",
              "Url": "resolucionpregradoTCOMTO/"+facultad.Id.toString(),
              "Opciones": null
            },{
              "Id": 17,
              "Nombre": "HCP",
              "Url": "resolucionpregradoHCP/"+facultad.Id.toString(),
              "Opciones": null
            },{
              "Id": 18,
              "Nombre": "HCH",
              "Url": "resolucionpregradoHCH/"+facultad.Id.toString(),
              "Opciones": null
            }]
          },{
            "Id": 9,
            "Nombre": "Posgrado",
            "Url": "url_nivel_2",
            "Opciones": [{
              "Id": 16,
              "Nombre": "HCP",
              "Url": "resolucionposgradoHCP/"+facultad.Id.toString(),
              "Opciones": null
            },{
              "Id": 17,
              "Nombre": "HCH",
              "Url": "resolucionposgradoHCH/"+facultad.Id.toString(),
              "Opciones": null
            }]
          }]          
        }
        contenidoInscripcion.push(itemInscripcion);
        contenidoResolucion.push(itemResolucion);
      });
      $scope.actual = "";
      $scope.token_service = token_service;
      $scope.breadcrumb = [];
      $scope.menu_service = [{ //aqui va el servicio de el app de configuracion
        "Id": 6,
        "Nombre": "Generación resolución",
        "Url": "url_nivel_1",
        "Opciones": contenidoResolucion
      },{ //aqui va el servicio de el app de configuracion
        "Id": 1,
        "Nombre": "Inscripción docentes",
        "Url": "url_nivel_1",
        "Opciones": contenidoInscripcion
      },{ 
        "Id": 11,
        "Nombre": "Carga académica",
        "Url": "url_nivel_1",
        "Opciones": [{
          "Id": 12,
          "Nombre": "Pregrado",
          "Url": "url_nivel_2",
          "Opciones": [{
            "Id": 13,
            "Nombre": "Asignar carga",
            "Url": "cargaPregrado",
            "Opciones": null
          }]
        },{
          "Id": 14,
          "Nombre": "Posgrado",
          "Url": "url_nivel_2",
          "Opciones": [{
            "Id": 15,
            "Nombre": "Asignar carga",
            "Url": "url_nivel_3",
            "Opciones": null
          }]
        }]
      }];
      recorrerArbol($scope.menu_service, "");
    });

    var recorrerArbol = function(item, padre) {
      var padres = "";
      for (var i = 0; i < item.length; i++) {
        if (item[i].Opciones === null) {
          padres = padre + " , " + item[i].Nombre;
          paths.push({
            'path': item[i].Url,
            'padre': padres.split(",")
          });
        } else {
          recorrerArbol(item[i].Opciones, padre + "," + item[i].Nombre);
        }
      }
      return padres;
    };

    var paths = [];

    var update_url = function() {
      $scope.breadcrumb = [''];
      for (var i = 0; i < paths.length; i++) {
        if ($scope.actual === "/" + paths[i].path) {
          $scope.breadcrumb = paths[i].padre;
        } else if ('/' === $scope.actual) {
          $scope.breadcrumb = [''];
        }
      }
    }
    

    $scope.$on('$routeChangeStart', function(next, current) {
      $scope.actual = $location.path();
      update_url();
    });


    //Pendiente por definir json del menu
    (function($) {
      $(document).ready(function() {
        $('ul.dropdown-menu [data-toggle=dropdown]').on('click', function(event) {
          event.preventDefault();
          event.stopPropagation();
          $(this).parent().siblings().removeClass('open');
          $(this).parent().toggleClass('open');
        });
      });
    })(jQuery);
  });
