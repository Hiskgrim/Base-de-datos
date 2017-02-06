'use strict';

describe('Controller: InscripcionPregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('inscripcionDocentesApp'));

  var InscripcionPregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    InscripcionPregradoCtrl = $controller('InscripcionPregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(InscripcionPregradoCtrl.awesomeThings.length).toBe(3);
  });
});
