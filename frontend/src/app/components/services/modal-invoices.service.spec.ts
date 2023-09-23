import { TestBed } from '@angular/core/testing';

import { ModalInvoicesService } from './modal-invoices.service';

describe('ModalInvoicesService', () => {
  let service: ModalInvoicesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ModalInvoicesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
