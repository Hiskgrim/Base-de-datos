'use strict';

describe('Service: cargaService', function () {

  // load the service's module
  beforeEach(module('asignacionCargaApp'));

  // instantiate service
  var cargaService;
  beforeEach(inject(function (_cargaService_) {
    cargaService = _cargaService_;
  }));

  it('should do something', function () {
    expect(!!cargaService).toBe(true);
  });

});
