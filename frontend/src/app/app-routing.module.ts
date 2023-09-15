import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LoginComponent } from './components/seguridad/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { ListInvoicesComponent } from './list/list-invoices/list-invoices.component';
import { RegisterComponent } from './components/seguridad/register/register.component';
import { ListInvoicesDetailsComponent } from './list/list-invoices-details/list-invoices-details.component';
import { ListInvoicesTableComponent } from './list/list-invoices-table/list-invoices-table.component';
import { ListInvoicesTableTotalComponent } from './list/list-invoices-table-total/list-invoices-table-total.component';

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'list_of_invoices', component: ListInvoicesComponent},
  {path: 'list_of_invoices_table', component: ListInvoicesTableComponent},
  {path: 'list_of_invoices_table_total', component: ListInvoicesTableTotalComponent},
  {path: 'list_and_details', component: ListInvoicesDetailsComponent},
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
  {path: '**', redirectTo: ''}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
