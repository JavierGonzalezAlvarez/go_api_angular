import { Component, OnInit, Input } from '@angular/core';

import { Invoices, serverResponse } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';
import { filter, pipe } from 'rxjs';

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

  displayedColumns: string[] = ['idheader', 'company', 'address', 'numberinvoice', 'datatime', 'createdat'];


  constructor(private listService: ServiceService) {
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
    console.log("to be done");
  }

  removeInvoice() {
    console.log("to be done");
  }

}
