'use strict';

/**
 * @ngdoc function
 * @name asignacionCargaApp.controller:CargaDocenteCtrl
 * @description
 * # CargaDocenteCtrl
 * Controller of the asignacionCargaApp
 */
angular.module('asignacionCargaApp')
  .controller('CargaDocenteCtrl', function (carga_request, $routeParams, $location) {

  	var self = this;

  	self.idDocente=parseInt($routeParams.IdDocente);
  	carga_request.getOne("docente",$routeParams.IdDocente).then(function(response){
      self.docente=response.data;
    });

    carga_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    carga_request.getAll("asignatura").then(function(response){
      self.asignaturas=response.data;
    });

    carga_request.getAll("grupo").then(function(response){
      self.grupos=response.data;
    });

    carga_request.getAll("docente_grupo").then(function(response){
      self.gruposdocentes=response.data;
    });

    self.asociarCarga = function(idGrupo){
      var docentegrupodatos = {
      	Horas: 6,
        IdDocente: {
          	Id: self.idDocente
        },
        IdGrupo: {
        	Id: idGrupo
        }
      };

      carga_request.post("docente_grupo",docentegrupodatos).then(function(response){
	    carga_request.getAll("docente_grupo").then(function(response){
	      self.gruposdocentes=response.data;
	    });
      });
    };

    self.eliminarCarga = function(idGrupoDocente){
      carga_request.delete("docente_grupo",idGrupoDocente).then(function(response){
	    carga_request.getAll("docente_grupo").then(function(response){
	      self.gruposdocentes=response.data;
	    });
      });
    }

    self.asignaturaLibre = function(grupo){
    	var libre=true;
    	self.gruposdocentes.forEach(function(grupodocente){
    		if(grupo.Id==grupodocente.IdGrupo.Id){
    			libre = false;
    		}
    	});
    	return libre;
    };

    self.regresar = function(){
      $location.url('/carga_pregrado');
    }

  });
