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
import { MatSnackBar } from '@angular/material/snack-bar';


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
  private snackBar = inject(MatSnackBar);

  prompt: string = '';
  email: string = '';
  is_loading: boolean = false;
  requestId: string = '';

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

  openSnackBar(msg: string): void {
    this.snackBar.open(
      msg,
      'Close', {
      horizontalPosition: "end",
      verticalPosition: "bottom",
    });
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
    this.is_loading = true;
    const url = environment.apiEndpoint + '/v1/generations';
    const body = { prompt: this.prompt, email: email };

    this.http.post<{ request_id: string }>(url, body).subscribe({
      next: (response) => {
        this.requestId = response.request_id;
        this.is_loading = true;
      },
      error: (err) => {
        console.error('Error in POST request:', err);
        this.is_loading = false;
        this.openSnackBar('Error in sending request');
      },
      complete: () => {
        console.log('POST request completed');
        this.is_loading = false;
      }
    });
  }

  pollGenerationStatus(): void {
    const url = environment.apiEndpoint + `/v1/generations/${this.requestId}`;

    const intervalId = setInterval(() => {
      if (!this.is_loading) {
        clearInterval(intervalId); // is_loading が false ならポーリング停止
        return;
      }

      this.http.get<{ status: string }>(url).subscribe({
        next: (response) => {
          console.log('Status response:', response.status);
          if (response.status !== 'pending') {
            this.is_loading = false; // 処理が完了したらポーリングを停止
            this.openSnackBar('Generation completed successfully!');
            clearInterval(intervalId); // ポーリング停止
          }
        },
        error: (err) => {
          console.error('Error fetching status:', err);
          this.is_loading = false;
          this.openSnackBar('Error fetching generation status');
          clearInterval(intervalId); // エラー発生時もポーリング停止
        }
      });
    }, 3000); // 3秒ごとにリクエストを送信
  }
}
