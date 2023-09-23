import { Injectable } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ModalInvoicesComponent } from '../modal-invoices/modal-invoices.component';

@Injectable({
  providedIn: 'root',
})
export class ModalInvoicesService {
  constructor(private dialog: MatDialog) {}

  openInvoiceModal(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;

    const dialogRef = this.dialog.open(ModalInvoicesComponent, dialogConfig);

    dialogRef.afterClosed().subscribe((invoiceData) => {
      if (invoiceData) {
        console.log('Invoice data from modal:', invoiceData);
      }
    });
  }
}
