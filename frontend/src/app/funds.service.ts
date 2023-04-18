import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { LoginSignUpService } from './loginsignuppage/loginsignuppage.service';
import { HttpClient, HttpResponse } from '@angular/common/http';


@Injectable({
  providedIn: 'root'
})
export class FundsService {
  private fundsSubject = new Subject<number>();

  constructor(
    private loginSignUpService: LoginSignUpService,
    private client: HttpClient
  ) { }

  async getFunds(): Promise<any> {
    const user = await this.loginSignUpService.claimData();
    const username = user.username;
    return this.client.post("/credentials/funds", { username }, { withCredentials: true }).toPromise();
  }
}
