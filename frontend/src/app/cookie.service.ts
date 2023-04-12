import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class CookieServices {
  constructor(private Cookies: CookieService) {
  }

  setNewCookie(cookieType: string, value: string) {
    const spoiledDate = new Date();
    spoiledDate.setDate(spoiledDate.getDate() + 1);
    this.Cookies.set(cookieType, value, {expires: spoiledDate});
  }

  getCookie(cookieType: string) {
    return this.Cookies.get(cookieType);
  }
  deleteCookie(cookieType: string) {
    this.Cookies.delete(cookieType)
  }
  checkForCookie(cookieType: string) {

  }
}
