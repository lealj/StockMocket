import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';
import { LoginsignuppageComponent } from './loginsignuppage.component';
import { LoginSignUpService } from '../../services/page-services/loginsignuppage.service';

describe('LoginsignuppageComponent', () => {
  let component: LoginsignuppageComponent;
  let fixture: ComponentFixture<LoginsignuppageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, FormsModule],
      declarations: [ LoginsignuppageComponent ],
      providers: [LoginSignUpService],
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginsignuppageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize login input boxes to empty strings', () => {
    expect(component.usernameL).toEqual('');
    expect(component.passwordL).toEqual('');
  });

  it('should initialize sign up input boxes to empty strings', () => {
    expect(component.usernameSU).toEqual('');
    expect(component.passwordSU).toEqual('');
  });
});
