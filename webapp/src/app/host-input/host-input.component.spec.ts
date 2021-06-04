import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HostInputComponent } from './host-input.component';

describe('HostInputComponent', () => {
  let component: HostInputComponent;
  let fixture: ComponentFixture<HostInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HostInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HostInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
