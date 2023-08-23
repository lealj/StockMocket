import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ShareService {

  constructor(private client: HttpClient) { }

  Buy(username: string, ticker: string, quantity: number): Promise <any> {
    const url = `userstock/buy`;
    const body = {
      username: username,
      ticker: ticker,
      shares: quantity
    };

    return firstValueFrom(this.client.post(url, body, {observe: 'response'}));
  }

  Sell(username: string, ticker: string, quantity: number): Promise <any>{
    const url = `userstock/sell`;
    const body = {
      username: username,
      ticker: ticker,
      shares: quantity
    };

    return firstValueFrom(this.client.post(url, body, {observe: 'response'}));
  }
}
