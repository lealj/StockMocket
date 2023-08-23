import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class LoginSignUpService {

  constructor(private client: HttpClient) { }

  AddOnLogin(username: string, password: string): Promise<any> {
    // accountInfo is passed to post request, and the http response is returned.
    const acctInfo = {username: username, password: password};
    return firstValueFrom(this.client.post("/credentials/login", acctInfo, {withCredentials: true, observe: 'response'}));
  }

  AddOnSignUp(username: string, password: string): Promise<any>
  {
    const acctInfo = { username: username, password: password };
    return firstValueFrom(this.client.post("/credentials/signup", acctInfo, {withCredentials: true, observe: 'response'}));
  }

  DeleteUNandPW(username: string): Promise<any>
  {
    const acctInfo = { username: username };
    return firstValueFrom(this.client.post("/credentials/delete", acctInfo, {withCredentials: true, observe: 'response'}));
  }

  verify(): Promise<boolean> {
    return firstValueFrom(this.client.get("/credentials/authorize", { withCredentials: true })).then(() => true)
      .catch(() => false);
  }

  claimData(): Promise<any> {
    return firstValueFrom(this.client.get("/credentials/authorize", {withCredentials: true}))
  }

  async logout(): Promise<void> {
    await firstValueFrom(this.client.get("/credentials/logout",{ withCredentials: true }))
  }
}
