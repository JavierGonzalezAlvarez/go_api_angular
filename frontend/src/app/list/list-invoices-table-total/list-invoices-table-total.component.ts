
import { Component, OnInit, Input } from '@angular/core';

import { Invoices, serverResponse } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';
import { filter, pipe } from 'rxjs';

@Component({
  selector: 'app-list-invoices-table-total',
  templateUrl: './list-invoices-table-total.component.html',
  styleUrls: ['./list-invoices-table-total.component.scss']
})
export class ListInvoicesTableTotalComponent {

  @Input()
  invoices_child_table_total: Invoices[] = [];

  public currentPage: number = 1;
  public size: number = 5;
  public totalItems: number = 0;
  public totalPages: number = 0;

  displayedColumns: string[] = ['idheader', 'company', 'address', 'numberinvoice', 'datatime', 'createdat'];


  constructor(private listService: ServiceService) {
    this.invoices_child_table_total = [];
  }

  ngOnInit(): void {
    this.loadPage(this.currentPage, this.size);
  }

  get_page_items(currentPage: number, totalItems: number, response: any []) {
    // we slice the results object depending on the elements i want per page to be shown
    const startIndex = (currentPage - 1) * this.size;
    const endIndex = startIndex + this.size;

    return response[0].results.slice(startIndex, endIndex);
  }

  loadPage( currentPage: number, totalItems: number ) {
    this.listService.get_all_headers_table_total(totalItems )
      .pipe(
        filter( (response: any[]) => response.length > 0 )
      )
      .subscribe( response => {
        console.log(response)
        // i get an array ob objects form the back
        this.currentPage = currentPage;
        this.totalItems = response[0].totalCount;
        this.totalPages = Math.ceil(this.totalItems / this.size);
        console.log(this.currentPage, this.totalItems, this.size);


        //this.invoices_child_table_total = response[0].results;
        this.invoices_child_table_total = this.get_page_items(this.currentPage, this.totalItems, response);
      })
  }

  onPageChange(page: number, size: number) {
    this.currentPage = page;
    this.loadPage(page, size);
  }


  addInvoice() {
    console.log("add invoice to be done");
  }

  removeInvoice() {
    console.log("remove invoice to be done");
  }

}
