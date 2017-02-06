'use strict';

describe('Controller: CargaPregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('asignacionCargaApp'));

  var CargaPregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargaPregradoCtrl = $controller('CargaPregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargaPregradoCtrl.awesomeThings.length).toBe(3);
  });
});
