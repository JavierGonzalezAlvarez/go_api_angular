import { Component, Input, OnInit } from '@angular/core';
import { SeguridadService } from '../seguridad.service';

@Component({
  selector: 'app-autorizado',
  templateUrl: './autorizado.component.html',
  styleUrls: ['./autorizado.component.scss']
})
export class AutorizadoComponent  implements OnInit{

  constructor(private seguridadService: SeguridadService) { }

  //@Input()
  //rol!: string;

  @Input()
  role!: string;

  ngOnInit(): void {
  }

  estaAutorizado(): boolean {
    if (this.role) {
      return (
        this.seguridadService.estaLogueado() && // Check if the user is logged in
        this.seguridadService.getRol() === this.role // Check if the user's role matches the expected role
      );
    } else {
      return this.seguridadService.estaLogueado(); // Check if the user is logged in without considering the role
    }
  }

/*
  estaAutorizado(): boolean {
    if (this.role) {
      console.log("there is a role: user")
      console.log("está logeado: ", this.seguridadService.estaLogueado())
      return this.seguridadService.estaLogueado();
      //return this.seguridadService.getRol() === "user"; //this.rol;
    } else {
      console.log("there is no role: ", this.role)
      console.log("está logeado: ", this.seguridadService.estaLogueado())
      return this.seguridadService.estaLogueado();
    }
  }
*/



}
