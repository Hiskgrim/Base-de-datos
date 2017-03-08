'use strict';

describe('Controller: ResolucionpregradooCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionpregradooCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionpregradooCtrl = $controller('ResolucionpregradooCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionpregradooCtrl.awesomeThings.length).toBe(3);
  });
});
