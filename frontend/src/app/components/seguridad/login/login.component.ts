import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SeguridadService } from '../seguridad.service';
import { parsearErroresAPI } from '../../utils/utils';
import { userCredentials } from '../seguridad';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

   constructor(private seguridadService: SeguridadService,
    private router: Router) { }

  errores: string[] = [];

  ngOnInit(): void {
  }

  login(credenciales: userCredentials) {
    this.seguridadService.login(credenciales)
      .subscribe(
        respuesta => {
          this.seguridadService.saveToken(respuesta);
          this.router.navigate(['/']);
      },
        errores => this.errores = parsearErroresAPI(errores));
  }

}
