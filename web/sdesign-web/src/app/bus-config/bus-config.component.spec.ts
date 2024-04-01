import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusConfigComponent } from './bus-config.component';

describe('BusConfigComponent', () => {
  let component: BusConfigComponent;
  let fixture: ComponentFixture<BusConfigComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BusConfigComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BusConfigComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
