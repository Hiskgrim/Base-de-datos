'use strict';

describe('Controller: InscripcionPosgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('inscripcionDocentesApp'));

  var InscripcionPosgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    InscripcionPosgradoCtrl = $controller('InscripcionPosgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(InscripcionPosgradoCtrl.awesomeThings.length).toBe(3);
  });
});
