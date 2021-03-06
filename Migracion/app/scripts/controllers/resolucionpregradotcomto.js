'use strict';

/**
 * @ngdoc function
 * @name cdveApp.controller:ResolucionpregradoCtrl
 * @description
 * # ResolucionpregradoCtrl
 * Controller of the cdveApp
 */
angular.module('cdveApp')
  .controller('ResolucionpregradotcomtoCtrl', function (vinculacion_request) {
    var self = this;

    vinculacion_request.getAll("proyecto_curricular").then(function(response){
      self.proyectos=response.data;
    });

    vinculacion_request.getAll("categoria").then(function(response){
      self.categorias=response.data;
    });

    vinculacion_request.getAll("dedicacion").then(function(response){
      self.dedicaciones=response.data;
      self.dedicacionesEspecificas=[];
      self.dedicaciones.forEach(function(dedicacion){
        if(dedicacion.Nombre=='MTO' || dedicacion.Nombre=='TCO'){
          self.dedicacionesEspecificas.push(dedicacion.Nombre);
        }
      });
    });

    vinculacion_request.getAll("ciudad").then(function(response){
      self.ciudades=response.data;
    });

    vinculacion_request.getAll("docente").then(function(response){
      self.docentes=response.data;
    });

    vinculacion_request.getAll("funcionario").then(function(response){
      self.funcionarios=response.data;
    });

    vinculacion_request.getAll("clasificacion").then(function(response){
      self.clasificaciones=response.data;
    });

    vinculacion_request.getAll("contrato_general").then(function(response){
      self.contratos=response.data;
    });

    vinculacion_request.getAll("contrato_especial").then(function(response){
      self.contratosEspeciales=response.data;
    });

    vinculacion_request.getAll("carga_academica").then(function(response){
      self.cargasAcademicas=response.data;
      self.datos = generarInformacion(self);
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.numero=resolucion["numero"];
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.preambulo=resolucion["preambulo"];
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.consideracion=resolucion["consideracion"];
    });
    
    $.getJSON("/resolucion.json", function(resolucion) {
        self.articulos=resolucion["articulos"];
    });
    
  self.agregarArticulo = function() {
    self.articulos.push({texto: '',
    paragrafo: null,
    asociado: false});  
  }

  self.eliminarArticulo = function(num) {
    self.articulos.splice(num,1);  
  }

  self.asociarTabla = function(num) {
    self.articulos.forEach(function(articulo){
      if(self.articulos[num]==articulo){
        articulo.asociado=true;
      }else{
        articulo.asociado=false;
      }
    })
  }

  self.generarResolucion = function() {
    var documento=getDocumento(self);
    pdfMake.createPdf(documento).getDataUrl(function(outDoc){
      document.getElementById('vistaPDF').src = outDoc;
    });
    $("#resolucion").show();
  }

  self.getNumeros = function(objeto) {
  	var numeros=[];
    if(objeto){
    	for(var i = 0; i<objeto.length; i++){
    		numeros.push(i);
    	}
    }
    return numeros;
  }

  });
