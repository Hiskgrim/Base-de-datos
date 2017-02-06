'use strict';

describe('Controller: ResolucionPosgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPosgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPosgradoCtrl = $controller('ResolucionPosgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPosgradoCtrl.awesomeThings.length).toBe(3);
  });
});
