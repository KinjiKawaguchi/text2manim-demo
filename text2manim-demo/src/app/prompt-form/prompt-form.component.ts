import { Component, inject } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MailAddrDialogComponent } from '../mail-addr-dialog/mail-addr-dialog.component';
import { MatDialog } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { SubmitButtonComponent } from '../submit-button/submit-button.component';

@Component({
  selector: 'app-prompt-form',
  standalone: true,
  imports: [
    CommonModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MailAddrDialogComponent,
    SubmitButtonComponent,
  ],
  templateUrl: './prompt-form.component.html',
  styleUrl: './prompt-form.component.css'
})
export class PromptFormComponent {
  private http = inject(HttpClient);
  private dialogService = inject(MatDialog);

  prompt: string = '';
  email: string = '';

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

    // ダイアログを開く
    const dialogRef = this.dialogService.open(MailAddrDialogComponent, {
      width: '60vh'
    });

    // ダイアログが閉じた後にPOSTリクエストを送信
    dialogRef.afterClosed().subscribe({
      next: (email) => {
        if (email) {
          this.sendPostRequest(email);
        }
      },
      error: (err) => {
        console.error('Error while closing dialog:', err);
      },
      complete: () => {
        console.log('Dialog closed successfully');
      }
    });
  }

  sendPostRequest(email: string): void {
    const url = environment.apiEndpoint + '/v1/generations';
    const body = { prompt: this.prompt, email: email };

    this.http.post(url, body).subscribe({
      next: (response) => {
        console.log('POST request successful:', response);
      },
      error: (err) => {
        console.error('Error in POST request:', err);
      },
      complete: () => {
        console.log('POST request completed');
      }
    });
  }
}
