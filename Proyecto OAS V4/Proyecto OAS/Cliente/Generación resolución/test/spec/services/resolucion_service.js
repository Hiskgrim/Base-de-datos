'use strict';

describe('Service: resolucionService', function () {

  // load the service's module
  beforeEach(module('generacionResolucionApp'));

  // instantiate service
  var resolucionService;
  beforeEach(inject(function (_resolucionService_) {
    resolucionService = _resolucionService_;
  }));

  it('should do something', function () {
    expect(!!resolucionService).toBe(true);
  });

});
