import { Component, OnInit } from '@angular/core';
import { ShareService } from '../share.service';
import { error } from 'cypress/types/jquery';
import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'


@Component({
  selector: 'app-buy-sell-button',
  templateUrl: './buy-sell-button.component.html',
  styleUrls: ['./buy-sell-button.component.scss']
})
export class BuySellButtonComponent implements OnInit {

  showBuyInput : boolean = false;
  showSellInput : boolean = false;
  quantity: number = 0;
  username = '';


  constructor(
    private shareAction: ShareService,
    private loginSignUpService: LoginSignUpService
    ){}

  ngOnInit(): void {}

  onBuyClick(): void {
    //when button is clicked we will show an input box for the quantity of shares
    this.showBuyInput = true;
    this.showSellInput = false;
  }

  async Buy(quantity: number){ //generalize to all tickers
    const userData = await this.loginSignUpService.claimData();
    this.username = userData.username;
    this.shareAction.Buy(this.username, "MSFT", quantity);
    this.showBuyInput = false; //hides input box after confirming order
  }

  onSellClick(){
    this.showSellInput = true;
    this.showBuyInput = false;
  }
  async Sell(quantity: number){ //generalize to all tickers
    const userData = await this.loginSignUpService.claimData();
    this.username = userData.username;
    this.shareAction.Sell(this.username, "MSFT", quantity);
    this.showSellInput = false; 
  }

}
