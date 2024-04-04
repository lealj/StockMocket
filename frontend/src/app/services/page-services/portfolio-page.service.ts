import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from "rxjs";
import { LoginSignUpService } from './loginsignuppage.service'

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
        return firstValueFrom(this.client.post("/portfoliovalue", accntInfo, {observe: 'response'}));
    }

    async getOwnedStocks(): Promise<any>
    {
        const userData = await this.loginSignUpService.claimData()
        const username = userData.username
        const accntInfo = { username: username}
        return firstValueFrom(this.client.post("/userstock/owned", accntInfo, {observe: 'response'}));
    }

    async getUserLogs(): Promise<any>
    {
        const userData = await this.loginSignUpService.claimData()
        const username = userData.username
        const accntInfo = { username: username}
        return firstValueFrom(this.client.post("/portfoliohistory", accntInfo, {observe: 'response'}));
    }
}