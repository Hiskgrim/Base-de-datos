'use strict';

describe('Controller: SolicitudPregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('solicitudConvocatoriaApp'));

  var SolicitudPregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    SolicitudPregradoCtrl = $controller('SolicitudPregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(SolicitudPregradoCtrl.awesomeThings.length).toBe(3);
  });
});
