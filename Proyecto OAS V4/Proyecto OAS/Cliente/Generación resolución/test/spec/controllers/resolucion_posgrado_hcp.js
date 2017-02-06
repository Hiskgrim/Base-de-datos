'use strict';

describe('Controller: ResolucionPosgradoHcpCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPosgradoHcpCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPosgradoHcpCtrl = $controller('ResolucionPosgradoHcpCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPosgradoHcpCtrl.awesomeThings.length).toBe(3);
  });
});
