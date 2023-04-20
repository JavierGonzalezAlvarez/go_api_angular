import { Component, Input, OnInit } from '@angular/core';

import { Invoices } from 'src/app/interfaces/invoices';

@Component({
  selector: 'app-list-general',
  templateUrl: './list-general.component.html',
  styleUrls: ['./list-general.component.scss']
})
export class ListGeneralComponent implements OnInit{

  @Input()
  list_invoices_general: Invoices[] = [];

  constructor() {
   }

   ngOnInit(): void {
  }

}
