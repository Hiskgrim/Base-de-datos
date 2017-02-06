'use strict';

describe('Controller: CargaPosgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('asignacionCargaApp'));

  var CargaPosgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargaPosgradoCtrl = $controller('CargaPosgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargaPosgradoCtrl.awesomeThings.length).toBe(3);
  });
});
