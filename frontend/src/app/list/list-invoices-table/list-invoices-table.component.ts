import { Component, OnInit, Input } from '@angular/core';
import { Observable } from 'rxjs';
import { Invoices, serverResponse } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';
import { filter, pipe } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment.local';
import { ModalInvoicesService } from 'src/app/components/services/modal-invoices.service';

@Component({
  selector: 'app-list-invoices-table',
  templateUrl: './list-invoices-table.component.html',
  styleUrls: ['./list-invoices-table.component.scss']
})
export class ListInvoicesTableComponent implements OnInit {

  @Input()
  invoices_child_table: Invoices[] = [];

  public currentPage: number = 1;
  public pageSize: number = 5;
  public totalItems: number = 0;
  public totalPages: number = 0;

  private apiURL = environment.urlBackend;

  displayedColumns: string[] = ['idheader', 'company', 'address', 'numberinvoice', 'datatime', 'createdat'];


  constructor(
    private listService: ServiceService,
    private httpClient: HttpClient,
    private modalService: ModalInvoicesService,
  ) {
    this.invoices_child_table = [];
  }

  ngOnInit(): void {
    this.loadPage(this.currentPage, this.totalItems);
  }

  loadPage( limitOfResults: number, totalItems: number ) {
    this.listService.get_all_headers_table( limitOfResults, totalItems )
      .pipe(
        filter( (invoices: any[]) => invoices.length > 0 )
      )
      .subscribe( invoices => {
        console.log(invoices)
        this.invoices_child_table = invoices
      })
  }

  onPageChange(page: number, size: number) {
    this.currentPage = page;
    this.loadPage(page, size);
  }

  addInvoice() {

    this.modalService.openInvoiceModal();

    const invoiceData = {
      "companyname": "jga_test112",
      "address": "cl none",
      "numberinvoice": 12
    };

    this.listService.createInvoice(invoiceData).subscribe(
      (response: any) => {
        console.log('Invoice header created successfully:', response);
      },
      (error) => {
        console.error('Error creating invoice:', error);
      }
    );
  }

    removeInvoice() {
      console.log("remove invoice to be done")
    }

}
