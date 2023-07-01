import { Component, Input, OnInit } from '@angular/core';
import { SeguridadService } from '../seguridad.service';

@Component({
  selector: 'app-autorizado',
  templateUrl: './autorizado.component.html',
  styleUrls: ['./autorizado.component.scss']
})
export class AutorizadoComponent  implements OnInit{

  constructor(private seguridadService: SeguridadService) { }

  @Input()
  rol!: string;

  ngOnInit(): void {

}

  estaAutorizado(): boolean {
    if (this.rol) {
      return this.seguridadService.getRol() === this.rol;
    } else {
      return this.seguridadService.estaLogueado();
    }
  }

}
