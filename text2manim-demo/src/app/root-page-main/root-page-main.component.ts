import { Component } from '@angular/core';
import { PromptFormComponent } from '../prompt-form/prompt-form.component';
import { GeneratedVideoPlayerComponent } from '../generated-video-player/generated-video-player.component';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-root-page-main',
  standalone: true,
  imports: [
    CommonModule,
    PromptFormComponent,
    GeneratedVideoPlayerComponent,
    MatButtonModule,
  ],
  templateUrl: './root-page-main.component.html',
  styleUrl: './root-page-main.component.css'
})
export class RootPageMainComponent {
  videoUrl: string = 'https://storage.googleapis.com/text2manim/videos/vid_1729077102307881077.mp4';
  // videoUrl: string = '';

  handleVideoUrlChange(videoUrl: string): void {
    this.videoUrl = videoUrl;
  }

  clearVideoUrl() {
    this.videoUrl = '';
  }
}
