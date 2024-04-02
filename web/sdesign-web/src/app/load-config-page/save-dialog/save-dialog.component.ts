import { Component, Inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

export interface SaveDialogData {
  saveas: boolean,
  filename: string
}

@Component({
  selector: 'app-save-dialog',
  standalone: true,
  imports: [
    FormsModule,
    MatDialogModule,
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule
  ],
  templateUrl: './save-dialog.component.html',
  styleUrl: './save-dialog.component.scss'
})
export class SaveDialogComponent {
  constructor(
    public dialogRef: MatDialogRef<SaveDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: SaveDialogData
  ) {

  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
