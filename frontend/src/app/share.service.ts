import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ShareService {

  constructor(private client: HttpClient) { }

  getUser (): Promise <any> {
    //remember to use backend function
    const url = 'credentials/NAME OF FUNCTION IN BACKEND FOR GETTING USER';
    return this.client.get(url, {observe: 'response'}).toPromise();
  }

  Buy(username: string, ticker: string, quantity: number): Promise <any> {
    const url = `userstock/${username}/buy`;
    const body = {
      ticker: ticker,
      shares: quantity
    };

    return this.client.post(url, body, {observe: 'response'}).toPromise();
  }

  Sell(username: string, ticker: string, quantity: number): Promise <any>{
    const url = `userstock/${username}/sell`;
    const body = {
      ticker: ticker,
      shares: quantity
    };

    return this.client.post(url, body, {observe: 'response'}).toPromise();
  }
}
