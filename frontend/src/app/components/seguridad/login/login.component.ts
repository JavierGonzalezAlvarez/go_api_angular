import { Component, OnInit} from '@angular/core';
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
    private router: Router,
    //private cdr: ChangeDetectorRef
  ) { }


  private readonly campoRol = 'role';
  role: string = "";
  errores: string[] = [];

  handleAPIErrors(errores: any) {
    this.errores = parsearErroresAPI(errores);
  }

  ngOnInit(): void {
  }


  get_role() {
      return this.seguridadService.getFieldJWT(this.campoRol);
  }

  login(credenciales: userCredentials) {
    this.seguridadService.login(credenciales)
      .subscribe(
        (respuesta: any) => {
          console.log('Login successfully');
          console.log('API Response:', respuesta); // Log the API response

          this.seguridadService.saveToken(respuesta)

          console.log("role: ", respuesta.role)

          // get role from JWT
          this.role = this.get_role();
          console.log("Role in JWT:", this.role)

          this.router.navigate(['/']);
        },

      errores => {
        this.handleAPIErrors(errores)
        console.log('API Errors:', this.errores)
      }
    );
  }

}
