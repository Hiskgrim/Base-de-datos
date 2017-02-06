'use strict';

describe('Service: solicitudService', function () {

  // load the service's module
  beforeEach(module('solicitudConvocatoriaApp'));

  // instantiate service
  var solicitudService;
  beforeEach(inject(function (_solicitudService_) {
    solicitudService = _solicitudService_;
  }));

  it('should do something', function () {
    expect(!!solicitudService).toBe(true);
  });

});
