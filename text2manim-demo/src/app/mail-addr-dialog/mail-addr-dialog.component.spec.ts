import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MailAddrDialogComponent } from './mail-addr-dialog.component';

describe('MailAddrDialogComponent', () => {
  let component: MailAddrDialogComponent;
  let fixture: ComponentFixture<MailAddrDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MailAddrDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MailAddrDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
