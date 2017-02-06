'use strict';

describe('Controller: SolicitudPosgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('solicitudConvocatoriaApp'));

  var SolicitudPosgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    SolicitudPosgradoCtrl = $controller('SolicitudPosgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(SolicitudPosgradoCtrl.awesomeThings.length).toBe(3);
  });
});
