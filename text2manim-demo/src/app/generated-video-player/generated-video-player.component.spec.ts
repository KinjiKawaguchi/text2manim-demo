import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GeneratedVideoPlayerComponent } from './generated-video-player.component';

describe('GeneratedVideoPlayerComponent', () => {
  let component: GeneratedVideoPlayerComponent;
  let fixture: ComponentFixture<GeneratedVideoPlayerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GeneratedVideoPlayerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GeneratedVideoPlayerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
