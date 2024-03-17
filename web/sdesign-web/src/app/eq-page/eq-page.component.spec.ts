import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EqPageComponent } from './eq-page.component';

describe('EqPageComponent', () => {
  let component: EqPageComponent;
  let fixture: ComponentFixture<EqPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EqPageComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(EqPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
