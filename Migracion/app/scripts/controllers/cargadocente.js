'use strict';

/**
 * @ngdoc function
 * @name cdveApp.controller:CargadocenteCtrl
 * @description
 * # CargadocenteCtrl
 * Controller of the cdveApp
 */
angular.module('cdveApp')
  .controller('CargadocenteCtrl', function (vinculacion_request, $routeParams, $location) {
    var self = this;

  	self.idDocente=parseInt($routeParams.IdDocente);
  	vinculacion_request.getOne("docente",$routeParams.IdDocente).then(function(response){
      self.docente=response.data;
    });

    vinculacion_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    vinculacion_request.getAll("asignatura").then(function(response){
      self.asignaturas=response.data;
    });

    vinculacion_request.getAll("grupo").then(function(response){
      self.grupos=response.data;
    });

    vinculacion_request.getAll("docente_grupo").then(function(response){
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

      vinculacion_request.post("docente_grupo",docentegrupodatos).then(function(response){
	    vinculacion_request.getAll("docente_grupo").then(function(response){
	      self.gruposdocentes=response.data;
	    });
      });
    };

    self.eliminarCarga = function(idGrupoDocente){
      vinculacion_request.delete("docente_grupo",idGrupoDocente).then(function(response){
	    vinculacion_request.getAll("docente_grupo").then(function(response){
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

    self.generarFormato = function() {
    var documento=getDocumento(self);
      pdfMake.createPdf(documento).getDataUrl(function(outDoc){
        document.getElementById('vistaPDF').src = outDoc;
      });
      $("#formato").show();
    }
  });
