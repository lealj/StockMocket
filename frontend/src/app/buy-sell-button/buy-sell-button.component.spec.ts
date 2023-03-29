import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BuySellButtonComponent } from './buy-sell-button.component';

describe('BuySellButtonComponent', () => {
  let component: BuySellButtonComponent;
  let fixture: ComponentFixture<BuySellButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BuySellButtonComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BuySellButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
