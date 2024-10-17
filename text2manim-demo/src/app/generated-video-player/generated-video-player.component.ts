import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-generated-video-player',
  standalone: true,
  imports: [],
  templateUrl: './generated-video-player.component.html',
  styleUrl: './generated-video-player.component.css'
})
export class GeneratedVideoPlayerComponent {
  @Input() videoUrl: string = '';
}
