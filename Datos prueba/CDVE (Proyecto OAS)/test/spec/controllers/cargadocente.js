'use strict';

describe('Controller: CargadocenteCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var CargadocenteCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargadocenteCtrl = $controller('CargadocenteCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargadocenteCtrl.awesomeThings.length).toBe(3);
  });
});
