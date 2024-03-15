import { TestBed } from '@angular/core/testing';

import { BackendWsService } from './backend-ws.service';

describe('BackendWsService', () => {
  let service: BackendWsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BackendWsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
