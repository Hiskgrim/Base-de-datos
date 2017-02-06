'use strict';

describe('Controller: ResolucionPosgradoTcoMtoCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPosgradoTcoMtoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPosgradoTcoMtoCtrl = $controller('ResolucionPosgradoTcoMtoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPosgradoTcoMtoCtrl.awesomeThings.length).toBe(3);
  });
});
