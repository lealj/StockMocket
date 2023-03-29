import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ShareService {

  constructor(private client: HttpClient) { }

  Buy(username: string, ticker: string, quantity: number): Promise <any> {
    const url = `userstock/${username}`;
    const body = {
      ticker: ticker,
      shares: quantity
    };

    return this.client.post(url, body, {observe: 'response'}).toPromise();
  }

  Sell(username: string, ticker: string, quantity: number): Promise <any>{
    const url = `userstock/${username}`;
    const body = {
      ticker: ticker,
      shares: quantity
    };

    return this.client.post(url, body, {observe: 'response'}).toPromise();
  }
}
