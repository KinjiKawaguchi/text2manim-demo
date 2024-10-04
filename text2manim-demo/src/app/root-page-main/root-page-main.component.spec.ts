import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RootPageMainComponent } from './root-page-main.component';

describe('RootPageMainComponent', () => {
  let component: RootPageMainComponent;
  let fixture: ComponentFixture<RootPageMainComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RootPageMainComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RootPageMainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
