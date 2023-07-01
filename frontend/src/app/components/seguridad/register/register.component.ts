import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { userCredentials } from '../seguridad';
import { SeguridadService } from '../seguridad.service';
import { parsearErroresAPI } from '../../utils/utils';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  constructor(private seguridadService: SeguridadService,
    private router: Router) { }

  ngOnInit(): void {

  }

  errores: string[] = [];

  register(credenciales: userCredentials) {
    this.seguridadService.register(credenciales)
      .subscribe(
        respuesta => {
          this.seguridadService.saveToken(respuesta);
          this.router.navigate(['/']);
      },
         errores => this.errores = parsearErroresAPI(errores));
  }
}
