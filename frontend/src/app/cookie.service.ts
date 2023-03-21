import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class CookieServices {
  constructor(private Cookies: CookieService) {
  }

  setNewCookie(username: string, value: string) {
    this.Cookies.set(username, value);
  }

  getCookie(username: string) {
    return this.Cookies.get(username);
  }
}
