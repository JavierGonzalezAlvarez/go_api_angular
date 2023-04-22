import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from "@angular/material/icon"
import { MatButtonModule } from "@angular/material/button"

import {MatTableModule} from '@angular/material/table';
import {CdkTableModule} from '@angular/cdk/table';

@NgModule({
  declarations: [],
  imports: [
    CommonModule
  ],
  exports: [
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    CdkTableModule,
    MatTableModule,
  ]
})
export class MaterialModule { }
