import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'
import { firstValueFrom } from 'rxjs';

@Injectable({
    providedIn: 'root'
})

export class StocksPageService {
    //construct
    constructor(
        private client: HttpClient,
        private loginSignUpService: LoginSignUpService 
    ) {}
    // async functions
    async getStocksList(): Promise<any> 
    {
        return firstValueFrom(this.client.get("/stocksdata", {observe: 'response'}))
    }
}