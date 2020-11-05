import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { KeyboardConfigComponent } from './keyboard-config.component';

describe('KeyboardConfigComponent', () => {
  let component: KeyboardConfigComponent;
  let fixture: ComponentFixture<KeyboardConfigComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyboardConfigComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyboardConfigComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
