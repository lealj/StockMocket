import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import {Observable, observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})

export class LoginSignUpService {

  constructor(private client: HttpClient) { }

  AddOnLogin(username: string, password: string): Promise<any> {
    // accountInfo is passed to post request, and the http response is returned.
    const acctInfo = {username: username, password: password};
    return this.client.post("/credentials/login", acctInfo, {withCredentials: true, observe: 'response'}).toPromise();
  }

  AddOnSignUp(username: string, password: string): Promise<any>
  {
    const acctInfo = { username: username, password: password };
    return this.client.post("/credentials/signup", acctInfo, {withCredentials: true, observe: 'response'}).toPromise();
  }

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



}
