'use strict';

describe('Controller: ResolucionPregradoTcoMtoCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPregradoTcoMtoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPregradoTcoMtoCtrl = $controller('ResolucionPregradoTcoMtoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPregradoTcoMtoCtrl.awesomeThings.length).toBe(3);
  });
});
