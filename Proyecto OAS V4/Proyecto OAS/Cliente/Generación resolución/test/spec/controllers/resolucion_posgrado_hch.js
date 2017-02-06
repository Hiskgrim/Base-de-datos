'use strict';

describe('Controller: ResolucionPosgradoHchCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPosgradoHchCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPosgradoHchCtrl = $controller('ResolucionPosgradoHchCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPosgradoHchCtrl.awesomeThings.length).toBe(3);
  });
});
