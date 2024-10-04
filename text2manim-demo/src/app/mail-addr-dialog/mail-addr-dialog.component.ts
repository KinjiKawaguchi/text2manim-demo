import { Component, inject, model } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule, MatDialogTitle, MatDialogActions, MatDialogClose, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-mail-addr-dialog',
  standalone: true,
  imports: [
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatDialogModule,
    FormsModule,
    MatDialogTitle,
    MatDialogActions,
  ],
  templateUrl: './mail-addr-dialog.component.html',
  styleUrl: './mail-addr-dialog.component.css'
})
export class MailAddrDialogComponent {
  private readonly dialogRef: MatDialogRef<MailAddrDialogComponent> = inject(MatDialogRef<MailAddrDialogComponent>);
  email: string = '';

  onClickClose(): void {
    this.dialogRef.close();
  }

  onSubmitEmail(): void {
    this.dialogRef.close(this.email);
  }

  isEmailValid(): boolean {
    return this.email.length > 0 && this.email.match(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/) !== null;
  }
}
