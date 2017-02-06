'use strict';

describe('Controller: ResolucionPregradoHchCtrl', function () {

  // load the controller's module
  beforeEach(module('generacionResolucionApp'));

  var ResolucionPregradoHchCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionPregradoHchCtrl = $controller('ResolucionPregradoHchCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionPregradoHchCtrl.awesomeThings.length).toBe(3);
  });
});
