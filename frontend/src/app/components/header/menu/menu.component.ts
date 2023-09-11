import { Component, OnInit } from '@angular/core';
import { SeguridadService } from '../../seguridad/seguridad.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent implements OnInit {
menu: any;

  constructor(public seguridadService: SeguridadService) { }

  ngOnInit(): void {
  }

}
