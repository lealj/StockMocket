import { Component, OnInit } from '@angular/core';
import { ShareService } from '../share.service';
import { error } from 'cypress/types/jquery';

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


  constructor(private shareAction: ShareService){}
  ngOnInit(): void {
   this.shareAction.getUser().then((username: string) => {
    this.username = username;
   }).catch((error) => {
    console.error(error);
   });
  }

  onBuyClick(): void {
    //when button is clicked we will show an input box for the quantity of shares
    this.showBuyInput = true;
    this.showSellInput = false;
  }

  Buy(quantity: number){ //generalize to all tickers
    this.shareAction.Buy(this.username, "msft", quantity);
    this.showBuyInput = false; //hides input box after confirming order
  }

  onSellClick(){
    this.showSellInput = true;
    this.showBuyInput = false;
  }
  Sell(quantity: number){ //generalize to all tickers
    this.shareAction.Sell(this.username, "msft", quantity);
    this.showSellInput = false; 
  }

}
