'use strict';

describe('Controller: ResolucionpregradohcpCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionpregradohcpCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionpregradohcpCtrl = $controller('ResolucionpregradohcpCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionpregradohcpCtrl.awesomeThings.length).toBe(3);
  });
});
