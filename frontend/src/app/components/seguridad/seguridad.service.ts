import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.local';
import { userCredentials, responseAuthentication, user } from './seguridad';

@Injectable({
  providedIn: 'root'
})
export class SeguridadService {

  constructor(private httpClient: HttpClient) { }

  apiURL = environment.urlBackend + 'users'

  private readonly llaveToken = 'token';
  private readonly llaveExpiracion = 'token-expiracion';
  private readonly campoRol = 'role';

  obtenerUsuarios(pagina: number, recordsPorPagina: number): Observable<any> {
    let params = new HttpParams();
    params = params.append('page', pagina.toString());
    params = params.append('recordsPerPage', recordsPorPagina.toString());
    return this.httpClient.get<user[]>(`${this.apiURL}/list_users`,
      { observe: 'response', params })
  }

  hacerAdmin(usuarioId: string) {
    const headers = new HttpHeaders('Content-Type: application/json');
    return this.httpClient.post(`${this.apiURL}/doAdmin`, JSON.stringify(usuarioId), { headers });
  }

  removerAdmin(usuarioId: string) {
    const headers = new HttpHeaders('Content-Type: application/json');
    return this.httpClient.post(`${this.apiURL}/removeAdmin`, JSON.stringify(usuarioId), { headers });
  }

  estaLogueado(): boolean {
    const token = localStorage.getItem(this.llaveToken);
    if (!token) {
      return false;
    }

    //const expiracion: string | null | Date = localStorage.getItem(this.llaveExpiracion);
    //const expiracionFecha = new Date(expiracion);

    //if (expiracionFecha <= new Date()) {
    //  this.logout();
    //  return false;
    // }
    return true;
  }

  logout() {
    localStorage.removeItem(this.llaveToken);
    localStorage.removeItem(this.llaveExpiracion);
  }

  getRol(): string {
    return this.getFieldJWT(this.campoRol);
  }

  register(credenciales: userCredentials): Observable<responseAuthentication> {
    return this.httpClient.post<responseAuthentication>(this.apiURL + '/create', credenciales);
  }

  login(credenciales: userCredentials): Observable<responseAuthentication> {
    return this.httpClient.post<responseAuthentication>(this.apiURL + '/login', credenciales);
  }

  saveToken(respuestaAutenticacion: responseAuthentication) {
    localStorage.setItem(this.llaveToken, respuestaAutenticacion.token);
    localStorage.setItem(this.llaveExpiracion, respuestaAutenticacion.expiracion.toString());
  }

  // get field to show in Menu
  getFieldJWT(campo: string): string {
    const token = localStorage.getItem(this.llaveToken);
    if (!token) { return ''; }
    var dataToken = JSON.parse(atob(token.split('.')[1]));
    return dataToken[campo];
  }

  getToken() {
    return localStorage.getItem(this.llaveToken);
  }

}
