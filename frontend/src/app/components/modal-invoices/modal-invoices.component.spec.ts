import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModalInvoicesComponent } from './modal-invoices.component';

describe('ModalInvoicesComponent', () => {
  let component: ModalInvoicesComponent;
  let fixture: ComponentFixture<ModalInvoicesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ModalInvoicesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ModalInvoicesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
