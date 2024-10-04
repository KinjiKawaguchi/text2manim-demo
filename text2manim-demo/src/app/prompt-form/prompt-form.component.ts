import { Component } from '@angular/core';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-prompt-form',
  standalone: true,
  imports: [
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    FormsModule,
  ],
  templateUrl: './prompt-form.component.html',
  styleUrl: './prompt-form.component.css'
})
export class PromptFormComponent {
  prompt: string = '';

  onSubmit(): void { };
}
