import { Component, OnInit } from '@angular/core';

import { HttpResponse } from '@angular/common/http';
import { Invoices } from 'src/app/interfaces/invoices';
import { ServiceService } from '../service.service';

@Component({
  selector: 'app-list-invoices',
  templateUrl: './list-invoices.component.html',
  styleUrls: ['./list-invoices.component.scss']
})
export class ListInvoicesComponent implements OnInit {

  invoices: Invoices[] = [];

  constructor(private listService: ServiceService) {
    this.invoices = [];
   }

  ngOnInit(): void {
    this.loadData();
  }

  loadData() {
    this.listService.getAll().subscribe((inv: Invoices[]) => {
      //console.log(inv)
      this.invoices =  inv;
      console.log("invoices", this.invoices);
    })
  };

}
