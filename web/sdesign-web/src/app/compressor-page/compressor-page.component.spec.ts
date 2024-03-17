import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompressorPageComponent } from './compressor-page.component';

describe('CompressorPageComponent', () => {
  let component: CompressorPageComponent;
  let fixture: ComponentFixture<CompressorPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CompressorPageComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CompressorPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
