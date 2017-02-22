'use strict';

describe('Controller: ResolucionpregradohchCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionpregradohchCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionpregradohchCtrl = $controller('ResolucionpregradohchCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionpregradohchCtrl.awesomeThings.length).toBe(3);
  });
});
