import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { SeguridadService } from '../../seguridad/seguridad.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent implements OnInit {

  isAuthenticated = false;
  userRole = ''; // Set the user's role here

  constructor(
    public seguridadService: SeguridadService,
    private cdr: ChangeDetectorRef
  ) { }

  ngOnInit(): void {
  }

  /* Opcion: use ChangeDetectorRef, when condition is ok, it refresh values and we don't need to refresh browser
  ngAfterViewInit(): void {
    // Determine if the user is authenticated and set the user's role
    this.isAuthenticated = this.seguridadService.estaLogueado();
    this.userRole = this.seguridadService.getRol();

    this.cdr.detectChanges();
  }
  */

}
