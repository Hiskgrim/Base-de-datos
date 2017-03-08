'use strict';

/**
 * @ngdoc function
 * @name cdveApp.controller:ResolucionpregradoCtrl
 * @description
 * # ResolucionpregradoCtrl
 * Controller of the cdveApp
 */
angular.module('cdveApp')
  .controller('ResolucionpregradotcomtoCtrl', function (vinculacion_request,$routeParams) {
    var self = this;

    self.IdFacultad=parseInt($routeParams.IdFacultad); 

    vinculacion_request.getAll("proyecto_curricular","query=IdFacultad.id%3A"+self.IdFacultad+"%2CIdTipoProyectoCurricular.Id%3A1&limit=0").then(function(response){
      self.proyectos=response.data;
    });

    vinculacion_request.getAll("contrato_especial").then(function(response){
      self.contratosEspeciales=response.data;
    });

    vinculacion_request.getAll("contrato_especial","query=IdFacultad.id%3A"+self.IdFacultad+"&limit=0").then(function(response){
      var contratoEspecial=response.data;
      self.datosContratos=[];
      self.proyectos.forEach(function(proyecto){
        self.datosContratos[proyecto.Id]=[];
      });
      if(contratoEspecial!=null){
        contratoEspecial.forEach(function(contratoEspecial){
          self.getContratoEspecial(contratoEspecial);
        });
      }
    });

    self.getContratoEspecial = function(contratoEspecial){
      var datosContrato={
        Nombre: null,
        Cedula: null,
        Expedida: null,
        Categoria: null,
        Dedicacion: null,
        HorasSemanales: null,
        PeriodoVinculacion: null,
        ValorContrato: null,
        ProyectoCurricular: null
      }
      vinculacion_request.getOne("contrato_general",contratoEspecial.NumeroContratoGeneral.Id).then(function(response){
        var contratoGeneral=response.data;
        datosContrato.ValorContrato=contratoGeneral.ValorContrato;
        datosContrato.PeriodoVinculacion=contratoGeneral.Vigencia;
      });
      vinculacion_request.getOne("funcionario",contratoEspecial.IdFuncionario.Id).then(function(response){
          var funcionario=response.data;
          datosContrato.Nombre=funcionario.PrimerNombre+" "+funcionario.SegundoNombre+" "+funcionario.PrimerApellido+" "+funcionario.SegundoApellido;
          datosContrato.Cedula=funcionario.Documento;
          vinculacion_request.getOne("ciudad",funcionario.IdCiudad.Id).then(function(response){
            var ciudad=response.data;
            datosContrato.Expedida=ciudad.Nombre;
          });
          vinculacion_request.getAll("docente","query=IdFuncionario.Id%3A"+funcionario.Id.toString()).then(function(response){
            var docente=response.data[0];
            vinculacion_request.getAll("clasificacion","query=IdDocente.Id%3A"+docente.Id.toString()).then(function(response){
              var clasificacion=response.data[0];
              vinculacion_request.getOne("categoria",clasificacion.IdCategoria.Id).then(function(response){
                var categoria=response.data;
                datosContrato.Categoria=categoria.Nombre;
              });
              vinculacion_request.getOne("dedicacion",clasificacion.IdDedicacion.Id).then(function(response){
                var dedicacion=response.data;
                datosContrato.Dedicacion=dedicacion.Nombre;
              });
              vinculacion_request.getAll("carga_academica","query=IdDocente.Id%3A"+docente.Id.toString()).then(function(response){
                var cargaAcademica=response.data[0];
                datosContrato.HorasSemanales=cargaAcademica.HorasSemanales;
                self.datosContratos[cargaAcademica.IdProyectoCurricular.Id].push(datosContrato);
              });
            });
          });
        });
    };

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
