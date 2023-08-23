import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { CookieServices } from './cookie.service';

describe('CookieService', () => {
  let service: CookieServices;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [CookieServices]
    }).compileComponents();
  
    service = TestBed.inject(CookieServices);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
