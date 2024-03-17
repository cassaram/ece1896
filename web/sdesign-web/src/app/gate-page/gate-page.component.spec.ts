import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GatePageComponent } from './gate-page.component';

describe('GatePageComponent', () => {
  let component: GatePageComponent;
  let fixture: ComponentFixture<GatePageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GatePageComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GatePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
