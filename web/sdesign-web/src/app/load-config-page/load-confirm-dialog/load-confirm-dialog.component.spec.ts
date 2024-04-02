import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadConfirmDialogComponent } from './load-confirm-dialog.component';

describe('LoadConfirmDialogComponent', () => {
  let component: LoadConfirmDialogComponent;
  let fixture: ComponentFixture<LoadConfirmDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoadConfirmDialogComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(LoadConfirmDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
