import { Component, Output, EventEmitter, inject } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MailAddrDialogComponent } from '../mail-addr-dialog/mail-addr-dialog.component';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CookieService } from 'ngx-cookie-service';
import { GenerationService } from '../services/generation.service';
import { SubmitButtonComponent } from '../submit-button/submit-button.component';

export const MAX_LENGTH: number = 150;

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
  styleUrls: ['./prompt-form.component.css'],
  providers: [CookieService],
})
export class PromptFormComponent {
  private dialogService = inject(MatDialog);
  private snackBar = inject(MatSnackBar);
  private cookieService = inject(CookieService);
  private generationService = inject(GenerationService);

  prompt: string = '';
  email: string = '';
  is_loading: boolean = false;
  requestId: string = '';
  videoUrl: string = '';

  readonly maxLength: number = MAX_LENGTH;
  charCount: number = 0;

  @Output() videoUrlChange: EventEmitter<string> = new EventEmitter<string>();

  constructor() {
    const savedEmail = this.cookieService.get('email');
    if (savedEmail) {
      this.email = savedEmail;
    }
  }

  private setVideoUrl(videoUrl: string): void {
    console.log('Setting video URL:', videoUrl);
    this.videoUrl = videoUrl;
    this.videoUrlChange.emit(videoUrl);
  }

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
    this.snackBar.open(msg, 'Close', {
      horizontalPosition: 'end',
      verticalPosition: 'bottom',
    });
  }

  onSubmit(): void {
    if (this.isOverLimit()) {
      return;
    }

    if (this.email) {
      this.sendPostRequest(this.email);
    } else {
      const dialogRef = this.dialogService.open(MailAddrDialogComponent, {
        width: '60vh',
      });

      dialogRef.afterClosed().subscribe({
        next: (email) => {
          if (email) {
            this.cookieService.set('email', email);
            this.email = email;
            this.sendPostRequest(email);
          }
        },
        error: (err) => {
          console.error('Error while closing dialog:', err);
        },
      });
    }
  }

  sendPostRequest(email: string): void {
    this.is_loading = true;

    this.generationService.sendGenerationRequest(this.prompt, email).subscribe({
      next: (response) => {
        this.requestId = response.request_id;
        this.is_loading = true;
        this.prompt = '';
        this.pollGenerationStatus();
      },
      error: (err) => {
        console.error('Error in POST request:', err);
        this.is_loading = false;
        this.openSnackBar('Error in sending request');
      },
    });
  }

  pollGenerationStatus(): void {
    // Mock Status
    this.is_loading = true;
    this.requestId = 'test1';

    const intervalId = setInterval(() => {
      if (!this.is_loading) {
        clearInterval(intervalId);
        return;
      }

      this.generationService.getGenerationStatus(this.requestId).subscribe({
        next: (response) => {
          if (response.Status == 'STATUS_COMPLETED') {
            this.is_loading = false;
            console.log('Generation completed:', response);
            this.setVideoUrl(response.VideoUrl);
            this.openSnackBar('Generation completed successfully!');
            clearInterval(intervalId);
          } else if (response.Status == 'STATUS_FAILED') {
            this.is_loading = false;
            this.openSnackBar('Generation failed');
            clearInterval(intervalId);
          }
        },
        error: (err) => {
          console.error('Error fetching status:', err);
          this.is_loading = false;
          this.openSnackBar('Error fetching generation status');
          clearInterval(intervalId);
        },
      });
    }, 10000);
  }
}
