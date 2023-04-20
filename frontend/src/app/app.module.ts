import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/seguridad/login/login.component';
import { ProfileComponent } from './components/profile/profile.component';
import { RegisterComponent } from './components/register/register.component';
import { ListGeneralComponent } from './list/list-general/list-general.component';
import { ListInvoicesComponent } from './list/list-invoices/list-invoices.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './material/material.module';
import { MenuComponent } from './components/header/menu/menu.component';
import { TitleComponent } from './components/header/title/title.component';
import { SeguridadComponent } from './components/seguridad/seguridad.component';

import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http'
import { SeguridadInterceptorService } from './components/seguridad/seguridad-interceptor.service';
import { ListInvoicesDetailsComponent } from './list/list-invoices-details/list-invoices-details.component';


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
    SeguridadComponent,
    ListInvoicesDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    HttpClientModule
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
