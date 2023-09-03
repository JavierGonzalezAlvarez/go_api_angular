import { Component, EventEmitter, Input, OnInit, Output  } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { FormGroup } from '@angular/forms';

import { userCredentials } from '../seguridad';

@Component({
  selector: 'app-formulario-autenticacion',
  templateUrl: './formulario-autenticacion.component.html',
  styleUrls: ['./formulario-autenticacion.component.scss']
})

export class FormularioAutenticacionComponent implements OnInit {


  constructor(private formBuilder: FormBuilder) { }
  form!: FormGroup;

  @Input()
  errores: string[] = [];
  @Input()
  accion!: string;
  @Output()

  onSubmit: EventEmitter<userCredentials> = new EventEmitter<userCredentials>();

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: ['', Validators.required],
      email: [
        '',
        {
          validators: [Validators.required, Validators.email],
        },
      ],
      password: [
        '',
        {
          validators: [Validators.required]
        }
      ]
    });
  }

  obtenerMensajeErrorEmail() {
    var campo = this.form.get('email');

    if (campo && campo.hasError('required')) {
      return 'El campo Email es requerido';
    }

    if (campo && campo.hasError('email')) {
      return 'El email no es válido';
    }

    return '';
  }

  /*
  obtenerMensajeErrorUserName() {
    var campo = this.form.get('username');

    if (campo && campo.hasError('required')) {
      return 'El campo Name es requerido';
    }

    if (campo && campo.hasError('username')) {
      return 'El name no es válido';
    }

    return '';
  }
  */
}
