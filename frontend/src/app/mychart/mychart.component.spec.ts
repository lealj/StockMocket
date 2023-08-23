import { ComponentFixture, TestBed } from '@angular/core/testing';
import { BuySellButtonComponent } from '../buy-sell-button/buy-sell-button.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MychartComponent } from './mychart.component';

describe('MychartComponent', () => {
  let component: MychartComponent;
  let fixture: ComponentFixture<MychartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      declarations: [ MychartComponent, BuySellButtonComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MychartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
