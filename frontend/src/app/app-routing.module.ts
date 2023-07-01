import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LoginComponent } from './components/seguridad/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { ListInvoicesComponent } from './list/list-invoices/list-invoices.component';
import { RegisterComponent } from './components/seguridad/register/register.component';
import { ListInvoicesDetailsComponent } from './list/list-invoices-details/list-invoices-details.component';

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'list_of_invoices', component: ListInvoicesComponent},
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
