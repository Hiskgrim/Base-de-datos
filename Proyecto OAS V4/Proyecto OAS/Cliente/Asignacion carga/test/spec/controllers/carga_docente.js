'use strict';

describe('Controller: CargaDocenteCtrl', function () {

  // load the controller's module
  beforeEach(module('asignacionCargaApp'));

  var CargaDocenteCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargaDocenteCtrl = $controller('CargaDocenteCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargaDocenteCtrl.awesomeThings.length).toBe(3);
  });
});
