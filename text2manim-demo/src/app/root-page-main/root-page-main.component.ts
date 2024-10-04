import { Component } from '@angular/core';
import { PromptFormComponent } from '../prompt-form/prompt-form.component';

@Component({
  selector: 'app-root-page-main',
  standalone: true,
  imports: [PromptFormComponent],
  templateUrl: './root-page-main.component.html',
  styleUrl: './root-page-main.component.css'
})
export class RootPageMainComponent {

}
