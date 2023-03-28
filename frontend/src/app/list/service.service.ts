import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Invoices } from '../interfaces/invoices';
import { environment } from 'src/environments/environment.local';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  constructor(private http: HttpClient) { }

  private urlBackend = environment.urlBackend;

  public getAll(): Observable<any> {
    return this.http.get<Invoices[]>(this.urlBackend + 'get', {observe: 'response'});
  }

}
