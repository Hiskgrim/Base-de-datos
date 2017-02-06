'use strict';

describe('Controller: ResolucionPregradoHcpCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPregradoHcpCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPregradoHcpCtrl = $controller('ResolucionPregradoHcpCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPregradoHcpCtrl.awesomeThings.length).toBe(3);
  });
});
