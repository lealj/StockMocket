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

  constructor(private shareAction: ShareService){}
  ngOnInit(): void {
      
  }

  onBuyClick(): void {
    //when button is clicked we will show an input box for the quantity of shares
    this.showBuyInput = true;
  }

  Buy(quantity: number){
    this.shareAction.Buy();
  }

  onSellClick(){
    this.showSellInput = true;
  }
  Sell(quantity: number){
    this.shareAction.Sell();
  }

}
