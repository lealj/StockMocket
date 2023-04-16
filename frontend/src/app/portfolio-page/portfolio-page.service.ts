import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import {Observable, observable} from "rxjs";
import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'

@Injectable({
    providedIn: 'root'
})

export class PortfolioPageService {

    constructor(
      private client: HttpClient,
      private loginSignUpService: LoginSignUpService  
      ) {}
    
    async getPortfolioValue(): Promise<any> 
    {
        const userData = await this.loginSignUpService.claimData()
        const username = userData.username
        const accntInfo = { username: username}
        // backend recieves the username, so problem likely not here
        return this.client.post("/portfoliovalue", accntInfo, {observe: 'response'}).toPromise();
    }
}