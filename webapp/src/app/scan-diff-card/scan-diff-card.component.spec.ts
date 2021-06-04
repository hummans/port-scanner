import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ScanDiffCardComponent } from './scan-diff-card.component';

describe('ScanDiffCardComponent', () => {
  let component: ScanDiffCardComponent;
  let fixture: ComponentFixture<ScanDiffCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ScanDiffCardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ScanDiffCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
