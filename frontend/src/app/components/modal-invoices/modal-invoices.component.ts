
import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';


@Component({
  selector: 'app-modal-invoices',
  templateUrl: './modal-invoices.component.html',
  styleUrls: ['./modal-invoices.component.scss']
})
export class ModalInvoicesComponent {


  constructor(
    public dialogRef: MatDialogRef<ModalInvoicesComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {}

  onSubmit(invoiceData: any): void {
    // Process and validate the invoice data as needed
    // Emit the invoice data to the parent component
    this.dialogRef.close(invoiceData);
  }
}

