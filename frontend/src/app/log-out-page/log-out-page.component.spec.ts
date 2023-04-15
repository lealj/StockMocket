import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';
import { LogOutPageComponent } from './log-out-page.component';
import { LogOutPageService } from './log-out-page.service';

describe('LogOutPageComponent', () => {
  let component: LogOutPageComponent;
  let fixture: ComponentFixture<LogOutPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, FormsModule],
      declarations: [ LogOutPageComponent ],
      providers: [LogOutPageService],
    })
    .compileComponents();

    fixture = TestBed.createComponent(LogOutPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

});
