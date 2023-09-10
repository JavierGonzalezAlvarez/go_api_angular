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


  displayedColumns: string[] = ['idheader', 'company', 'address', 'numberinvoice', 'datatime', 'createdat'];

  constructor(private listService: ServiceService) {
    this.invoices_child_table = [];
  }

  ngOnInit(): void {
    this.loadPage(this.currentPage);
  }

  loadPage( page: number ) {
    this.listService.get_all_headers_table( page )
      .pipe(
        filter( (invoices: any[]) => invoices.length > 0 )
      )
      .subscribe( invoices => {
        console.log(invoices)
        this.currentPage = page;
        this.invoices_child_table = invoices
      })
  }

  /*
  loadData() {
    this.listService.getAll()
      .subscribe((invs: serverResponse["invoices"]) => {
        this.invoices_child_table = invs;
        console.log("invoices: ", this.invoices_child_table)
    });
  }
  */


  addInvoice() {
    console.log("to be done");
  }

  removeInvoice() {
    console.log("to be done");
  }

}
