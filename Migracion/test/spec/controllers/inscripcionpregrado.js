'use strict';

describe('Controller: InscripcionpregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var InscripcionpregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    InscripcionpregradoCtrl = $controller('InscripcionpregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(InscripcionpregradoCtrl.awesomeThings.length).toBe(3);
  });
});
