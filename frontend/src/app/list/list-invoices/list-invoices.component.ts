import { Component, Input, OnInit } from '@angular/core';

import { HttpResponse } from '@angular/common/http';
import { Invoices, serverResponse } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';

@Component({
  selector: 'app-list-invoices',
  templateUrl: './list-invoices.component.html',
  styleUrls: ['./list-invoices.component.scss']
})
export class ListInvoicesComponent implements OnInit {

  @Input()
  invoices_child: Invoices[] = [];

  constructor(private listService: ServiceService) {
    this.invoices_child = [];
   }

  ngOnInit(): void {
    this.loadData();
 
  }

  loadData() {
    this.listService.getAll().subscribe((invs: serverResponse["invoices"]) => {
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
