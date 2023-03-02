import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';
import { LoginsignuppageComponent } from './loginsignuppage.component';
import { LoginSignUpService } from './loginsignuppage.service';

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

  it('should initialize usernameL and passwordL to empty strings', () => {
    expect(component.usernameL).toEqual('');
    expect(component.passwordL).toEqual('');
  });

  it('should reset usernameL and passwordL to empty strings after calling AddOnLogin() method', (async() => {
    const loginService = TestBed.inject(LoginSignUpService); // create an instance of LoginSignUpService
    spyOn(loginService, 'AddOnLogin').and.returnValue(Promise.resolve());
    component.usernameL = 'testuser';
    component.passwordL = 'testpass';
    component.AttemptLogin();
    fixture.whenStable().then(() => {
      expect(component.usernameL).toEqual('');
      expect(component.passwordL).toEqual('');
    });
  }));
});
