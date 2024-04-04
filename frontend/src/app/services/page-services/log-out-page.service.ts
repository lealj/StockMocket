import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { firstValueFrom } from "rxjs";
import { LoginSignUpService } from './loginsignuppage.service'

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
    return firstValueFrom(this.client.post("/credentials/delete", acctInfo, {observe: 'response'}));
  }

  verify(): Promise<boolean> {
    return firstValueFrom(this.client.get("/credentials/authorize", { withCredentials: true })).then(() => true)
      .catch(() => false);
  }

  async logout(): Promise<void> {
    await firstValueFrom(this.client.get("/credentials/logout",{ withCredentials: true }));
  }

  async resetAccount(): Promise<any>
  {
    const userData = await this.loginSignUpService.claimData()
    const username = userData.username
    const accntInfo = { username: username}
    return firstValueFrom(this.client.post("/resetaccount", accntInfo, {observe: 'response'}));
  }

}
