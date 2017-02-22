'use strict';

describe('Controller: ResolucionposgradohcpCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionposgradohcpCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionposgradohcpCtrl = $controller('ResolucionposgradohcpCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionposgradohcpCtrl.awesomeThings.length).toBe(3);
  });
});
