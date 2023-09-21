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

  errores: string[] = [];

  handleAPIErrors(errores: any) {
    this.errores = parsearErroresAPI(errores);
  }

  ngOnInit(): void {

  }

  register(credenciales: userCredentials) {
    this.seguridadService.register(credenciales)
      .subscribe(
      (respuesta: any) => {
        console.log('Registration successfully');
        console.log('API Response:', respuesta); // Log the API response
        this.router.navigate(['/']);
      },
        errores => {
          this.handleAPIErrors(errores)
          console.log('API Errors:', this.errores)
        }
      );

  }
}
