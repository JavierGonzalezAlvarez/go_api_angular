import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/seguridad/login/login.component';
import { ProfileComponent } from './components/profile/profile.component';
import { ListGeneralComponent } from './list/list-general/list-general.component';
import { ListInvoicesComponent } from './list/list-invoices/list-invoices.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material/material.module';
import { MenuComponent } from './components/header/menu/menu.component';
import { TitleComponent } from './components/header/title/title.component';
import { RegisterComponent } from './components/seguridad/register/register.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http'
import { SeguridadInterceptorService } from './components/seguridad/seguridad-interceptor.service';
import { ListInvoicesDetailsComponent } from './list/list-invoices-details/list-invoices-details.component';
import { AutorizadoComponent } from './components/seguridad/autorizado/autorizado.component';
import { ValidatorsComponent } from './components/utils/validators/validators.component';
import { FormularioAutenticacionComponent } from './components/seguridad/formulario-autenticacion/formulario-autenticacion.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ListInvoicesTableComponent } from './list/list-invoices-table/list-invoices-table.component';
import { ListInvoicesTableTotalComponent } from './list/list-invoices-table-total/list-invoices-table-total.component';
import { ModalInvoicesComponent } from './components/modal-invoices/modal-invoices.component';
import { FormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    HomeComponent,
    LoginComponent,
    ProfileComponent,
    RegisterComponent,
    ListGeneralComponent,
    ListInvoicesComponent,
    MenuComponent,
    TitleComponent,
    ListInvoicesDetailsComponent,
    AutorizadoComponent,
    ValidatorsComponent,
    FormularioAutenticacionComponent,
    ListInvoicesTableComponent,
    ListInvoicesTableTotalComponent,
    ModalInvoicesComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule,
    MatDialogModule,
  ],
  providers: [
    // error handling in the response
    //{
    //  provide: HTTP_INTERCEPTORS,
    //  useClass: SeguridadInterceptorService,
    //  multi: true
    //}
],
  bootstrap: [AppComponent]
})
export class AppModule { }
