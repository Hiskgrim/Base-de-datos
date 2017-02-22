'use strict';

describe('Controller: ResolucionpregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var ResolucionpregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ResolucionpregradoCtrl = $controller('ResolucionpregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(ResolucionpregradoCtrl.awesomeThings.length).toBe(3);
  });
});
