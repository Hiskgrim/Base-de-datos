'use strict';

describe('Service: inscripcionService', function () {

  // load the service's module
  beforeEach(module('inscripcionDocentesApp'));

  // instantiate service
  var inscripcionService;
  beforeEach(inject(function (_inscripcionService_) {
    inscripcionService = _inscripcionService_;
  }));

  it('should do something', function () {
    expect(!!inscripcionService).toBe(true);
  });

});
