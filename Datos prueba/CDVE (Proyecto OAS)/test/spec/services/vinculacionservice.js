'use strict';

describe('Service: vinculacionService', function () {

  // load the service's module
  beforeEach(module('cdveApp'));

  // instantiate service
  var vinculacionService;
  beforeEach(inject(function (_vinculacionService_) {
    vinculacionService = _vinculacionService_;
  }));

  it('should do something', function () {
    expect(!!vinculacionService).toBe(true);
  });

});
