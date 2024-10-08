import { Component, Input, Output, EventEmitter } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-submit-button',
  standalone: true,
  imports: [
    MatIconModule,
  ],
  templateUrl: './submit-button.component.html',
  styleUrl: './submit-button.component.css'
})
export class SubmitButtonComponent {
  @Input() disabled: boolean = false;
  @Output() submit = new EventEmitter<void>();

  onClick(): void {
    this.submit.emit();
  }
}

