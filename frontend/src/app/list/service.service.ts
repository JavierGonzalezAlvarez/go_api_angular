import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable, map } from 'rxjs';

//interface
import { Invoices, serverResponse } from '../interfaces/invoices';
import { environment } from 'src/environments/environment.local';


@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  constructor(private HttpClient: HttpClient) { }

  private urlBackend = environment.urlBackend;

  //public getAll(): Observable<any> {
  //  return this.http.get<Invoices[]>(this.urlBackend + 'get', {observe: 'response'});
  //}


  public get_all_headers(limitOfResults=3): Observable<serverResponse["invoices"]> {
    return this.HttpClient.get<serverResponse["invoices"]>(this.urlBackend + 'get_all_header_invoices', {
      params: {
        limit: limitOfResults.toString()
      }
    });
  }

  //loaded data - list table
  public get_all_headers_table(limitOfResults=3, totalItems: number): Observable<serverResponse["invoices"]> {
    return this.HttpClient.get<serverResponse["invoices"]>(this.urlBackend + 'get_all_header_invoices', {
      params: {
        limitOfResults: limitOfResults,
        totalItems: totalItems
      }
    })
    .pipe(map( response => response));
  }

  public get_all_headers_table_total(totalItems: number): Observable<serverResponse["invoices"]> {
    return this.HttpClient.get<serverResponse["invoices"]>(this.urlBackend + 'get_all_header_invoices_total', {
      params: {
        totalItems: totalItems
      }
    })
    .pipe(map( response => response));
  }

  public send_new_invoice() {
    console.log("sending a post to create a new invoice")

  }

  public createInvoice(formData: any): Observable<{ message: string }> {
    const {companyname, address, numberinvoice} = formData;
    console.log("data in form: ", formData);
    return this.HttpClient.post<{ message: string }>(`${this.urlBackend}createOneHeader`, {
      companyname,
      address,
      numberinvoice,
    });
  }



}
