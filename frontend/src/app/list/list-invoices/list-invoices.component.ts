import { Component, Input, OnInit } from '@angular/core';

import { Invoices, serverResponse } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';
import { filter, pipe } from 'rxjs';

import {DataSource} from '@angular/cdk/collections';
import {Observable, ReplaySubject} from 'rxjs';

export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
}

const ELEMENT_DATA: PeriodicElement[] = [
  {position: 1, name: 'Hydrogen', weight: 1.0079, symbol: 'H'},
  {position: 2, name: 'Helium', weight: 4.0026, symbol: 'He'},
  {position: 3, name: 'Lithium', weight: 6.941, symbol: 'Li'},
  {position: 4, name: 'Beryllium', weight: 9.0122, symbol: 'Be'},
  {position: 5, name: 'Boron', weight: 10.811, symbol: 'B'},
  {position: 6, name: 'Carbon', weight: 12.0107, symbol: 'C'},
  {position: 7, name: 'Nitrogen', weight: 14.0067, symbol: 'N'},
  {position: 8, name: 'Oxygen', weight: 15.9994, symbol: 'O'},
  {position: 9, name: 'Fluorine', weight: 18.9984, symbol: 'F'},
  {position: 10, name: 'Neon', weight: 20.1797, symbol: 'Ne'},
];


@Component({
  selector: 'app-list-invoices',
  templateUrl: './list-invoices.component.html',
  styleUrls: ['./list-invoices.component.scss']
})
export class ListInvoicesComponent implements OnInit {

  @Input()
  invoices_child: Invoices[] = [];
  
  displayedColumns: string[] = ['idheader', 'company', 'address', 'numberinvoice', 'datatime', 'createdat'];
     
  public currentPage: number = 1;

  constructor(private listService: ServiceService) {
    this.invoices_child = [];
  }
 
  addInvoice() {
    console.log("to be done");
  }

  removeInvoice() {
    console.log("to be done");
  }

  ngOnInit(): void {
    this.loadPage(this.currentPage);
  }
  
  /*
  loadData( page: number ) {
    this.listService.getAll().subscribe((invs: serverResponse["invoices"]) => {
      this.invoices_child = invs;
      console.log("invoices: ", this.invoices_child)   
    });  
  }
 */

  loadPage( page: number ) {
    this.listService.loadPage( page )
      .pipe(
        filter( (invoices: any[]) => invoices.length > 0 )
      )
      .subscribe( invoices => {
        console.log(invoices)
        this.currentPage = page;
        this.invoices_child = invoices
      })
  }

  loadData() {
    this.listService.getAll()
      .subscribe((invs: serverResponse["invoices"]) => {
        this.invoices_child = invs;
        console.log("invoices: ", this.invoices_child)   
    }); 
  }

   /*
  loadData() {
    this.listService.getAll().subscribe((inv: Invoices[]) => {
      //console.log(inv)
      this.invoices =  inv;
      console.log("invoices", this.invoices);
    })
  };
  */
}

