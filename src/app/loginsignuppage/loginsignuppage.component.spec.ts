import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginsignuppageComponent } from './loginsignuppage.component';

describe('LoginsignuppageComponent', () => {
  let component: LoginsignuppageComponent;
  let fixture: ComponentFixture<LoginsignuppageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoginsignuppageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginsignuppageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
