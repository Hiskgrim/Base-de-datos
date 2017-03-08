'use strict';

describe('Controller: CargapregradoCtrl', function () {

  // load the controller's module
  beforeEach(module('cdveApp'));

  var CargapregradoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    CargapregradoCtrl = $controller('CargapregradoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(CargapregradoCtrl.awesomeThings.length).toBe(3);
  });
});
