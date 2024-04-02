import { Component, Inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialog, MatDialogModule } from '@angular/material/dialog';

export interface LoadConfirmDialogData {
  name: string,
  filename: string
}

@Component({
  selector: 'app-load-confirm-dialog',
  standalone: true,
  imports: [
    MatDialogModule,
    MatButtonModule
  ],
  templateUrl: './load-confirm-dialog.component.html',
  styleUrl: './load-confirm-dialog.component.scss'
})
export class LoadConfirmDialogComponent {
  constructor(@Inject(MAT_DIALOG_DATA) public data: LoadConfirmDialogData) {

  }
}
