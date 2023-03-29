import { Component, OnInit } from '@angular/core';
import { ShareService } from '../share.service';

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
   
  }

  onBuyClick(): void {
    //when button is clicked we will show an input box for the quantity of shares
    this.showBuyInput = true;
    this.showSellInput = false;
  }

  Buy(quantity: number){
    username: this.username;
    this.shareAction.Buy(this.username, "msft", quantity);
    this.showBuyInput = false; //hides input box after confirming order
  }

  onSellClick(){
    this.showSellInput = true;
    this.showBuyInput = false;
  }
  Sell(quantity: number){
    username: this.username;
    this.shareAction.Sell(this.username, "msft", quantity);
    this.showSellInput = false; 
  }

}
