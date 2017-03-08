'use strict';

describe('Controller: CargaposgradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var CargaposgradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargaposgradoCtrl = $controller('CargaposgradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargaposgradoCtrl.awesomeThings.length).toBe(3);
  });
});
