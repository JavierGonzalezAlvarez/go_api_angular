import { TestBed } from '@angular/core/testing';

import { SeguridadInterceptorService } from './seguridad-interceptor.service';

describe('SeguridadInterceptorService', () => {
  let service: SeguridadInterceptorService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SeguridadInterceptorService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
