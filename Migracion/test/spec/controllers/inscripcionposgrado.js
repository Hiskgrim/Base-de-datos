'use strict';

describe('Controller: InscripcionposgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var InscripcionposgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    InscripcionposgradoCtrl = $controller('InscripcionposgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(InscripcionposgradoCtrl.awesomeThings.length).toBe(3);
  });
});
