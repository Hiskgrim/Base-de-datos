'use strict';

describe('Controller: ResolucionposgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionposgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionposgradoCtrl = $controller('ResolucionposgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionposgradoCtrl.awesomeThings.length).toBe(3);
  });
});
