import { Component, inject } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { MailAddrDialogComponent } from '../mail-addr-dialog/mail-addr-dialog.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-prompt-form',
  standalone: true,
  imports: [
    CommonModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    FormsModule,
    MailAddrDialogComponent,
  ],
  templateUrl: './prompt-form.component.html',
  styleUrl: './prompt-form.component.css'
})
export class PromptFormComponent {
  private dialogService = inject(MatDialog);

  prompt: string = '';

  charCount: number = 0;
  maxLength: number = 150;

  updateCharCount(): void {
    this.charCount = this.prompt.length;
  }

  isOverLimit(): boolean {
    return this.charCount >= this.maxLength;
  }

  isNoInput(): boolean {
    return this.charCount === 0;
  }

  onSubmit(): void {
    if (this.isOverLimit()) {
      return;
    }

    // Open dialog
    const dialogRef = this.dialogService.open(MailAddrDialogComponent, {
      width: '60vh'
    })

    // Handle dialog result
    dialogRef.afterClosed().subscribe(email => {
      console.log(`Dialog result: ${email}`);
    });
  };
}
