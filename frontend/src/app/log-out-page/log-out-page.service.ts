import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import {Observable, observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})

export class LogOutPageService {

  constructor(private client: HttpClient) { }

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
