import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import {Observable, observable} from "rxjs";
import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'

@Injectable({
  providedIn: 'root'
})

export class LogOutPageService {

  constructor(
    private client: HttpClient,
    private loginSignUpService: LoginSignUpService  
    ) {}

  DeleteUNandPW(username: string): Promise<any>
  {
    const acctInfo = { username: username };
    return this.client.post("/credentials/delete", acctInfo, {observe: 'response'}).toPromise();
  }

  verify(): Promise<boolean> {
    return this.client.get("/credentials/authorize", { withCredentials: true }).toPromise().then(() => true)
      .catch(() => false);
  }

  async logout(): Promise<void> {
    await this.client.get("/credentials/logout",{ withCredentials: true }).toPromise()
  }

  async resetAccount(): Promise<any>
  {
    const userData = await this.loginSignUpService.claimData()
    const username = userData.username
    const accntInfo = { username: username}
    return this.client.post("/resetaccount", accntInfo, {observe: 'response'}).toPromise();
  }

}
