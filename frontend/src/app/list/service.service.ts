import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

//interface
import { Invoices, serverResponse } from '../interfaces/invoices';
import { environment } from 'src/environments/environment.local';


@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  constructor(private http: HttpClient) { }

  private urlBackend = environment.urlBackend;

  //public getAll(): Observable<any> {
  //  return this.http.get<Invoices[]>(this.urlBackend + 'get', {observe: 'response'});
  //}

  public getAll(limitOfResults=3): Observable<serverResponse["invoices"]> {
    return this.http.get<serverResponse["invoices"]>(this.urlBackend + 'get', {
      params: {
        limit: limitOfResults.toString()
      }
    });
  }

}
