import { ComponentFixture, TestBed, tick, fakeAsync } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';
import { BuySellButtonComponent } from './buy-sell-button.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';

describe('BuySellComponent', () => {
  let component: BuySellButtonComponent;
  let fixture: ComponentFixture<BuySellButtonComponent>;
  let buyButton: DebugElement;
  let sellButton: DebugElement;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, FormsModule],
      declarations: [ BuySellButtonComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BuySellButtonComponent);
    component = fixture.componentInstance;
    buyButton = fixture.debugElement.query(By.css('[data-cy="buy_button"]'));
    sellButton = fixture.debugElement.query(By.css('[data-cy="sell_button"]'));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should show the buy input box when buy button is clicked', fakeAsync(() => {
    buyButton.triggerEventHandler('click', null);
    tick();
    fixture.detectChanges();
    const buyInput = fixture.debugElement.query(By.css('#quantity'));
    expect(buyInput).toBeTruthy();
  }));

  it('should show the user name input box when buy button is clicked', fakeAsync(() => {
    sellButton.triggerEventHandler('click', null);
    tick();
    fixture.detectChanges();
    const sellInput = fixture.debugElement.query(By.css('#Username'));
    expect(sellInput).toBeTruthy();
  }));

  it('should show the sell input box when sell button is clicked', fakeAsync(() => {
    sellButton.triggerEventHandler('click', null);
    tick();
    fixture.detectChanges();
    const sellInput = fixture.debugElement.query(By.css('#quantity'));
    expect(sellInput).toBeTruthy();
  }));

  it('should show the user name input box when sell button is clicked', fakeAsync(() => {
    sellButton.triggerEventHandler('click', null);
    tick();
    fixture.detectChanges();
    const sellInput = fixture.debugElement.query(By.css('#Username'));
    expect(sellInput).toBeTruthy();
  }));
});
