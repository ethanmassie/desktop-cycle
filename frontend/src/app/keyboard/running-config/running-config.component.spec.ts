import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RunningConfigComponent } from './running-config.component';

describe('RunningConfigComponent', () => {
  let component: RunningConfigComponent;
  let fixture: ComponentFixture<RunningConfigComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RunningConfigComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RunningConfigComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
