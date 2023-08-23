import { Component, OnInit } from '@angular/core';
import { ShareService } from '../services/share.service';
import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'
import { Router, ActivatedRoute } from '@angular/router';

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
  errorMessageToPrint = '';
  ticker: string = '';


  constructor(
    private shareAction: ShareService,
    private loginSignUpService: LoginSignUpService,
    private router : Router,
    private route : ActivatedRoute
    )
    {
      this.route.paramMap.subscribe(params => {
        const tickerParam = params.get('ticker');
        this.ticker = tickerParam ?? '';
      });
    }

  ngOnInit(): void {}

  onBuyClick(): void {
    //when button is clicked we will show an input box for the quantity of shares
    this.showBuyInput = true;
    this.showSellInput = false;
  }

  async Buy(quantity: number){ //generalize to all tickers
    //make sure user is logged in before buying
    this.loginSignUpService.claimData().then((response) => {
      
    }).catch((error) => {
      console.log(error);
      if(error){
        this.errorMessageToPrint = "Please login first!"; //in the future this can be made a link
      }
    })

    const userData = await this.loginSignUpService.claimData();
    this.username = userData.username;
    this.shareAction.Buy(this.username, this.ticker, quantity).then((response) => { //make http request and wait for response, upon response send message to user
      if (response.status == 200) {
        this.errorMessageToPrint = "Shares successfully bought! updating funds...";
      }


      //to update the funds text every time the funds change the page is reloaded
      setTimeout(() => { 
        this.router.navigate(['.'], { relativeTo: this.route }).then(() => {
          location.reload();
        });
      }, 2000); // Delay of 2 seconds to read text

    }).catch((error) => {
        console.error(error);
        this.handleBuyError(error);
    });

    this.showBuyInput = false; //hides input box after confirming order
  }

  private handleBuyError(error: any): void {
    const errorStatusMessages: Record<number, string> = {
      400: "Username not found",
      401: "Ticker not found",
      402: "Share quantity is not in range 1-50",
      403: "Not enough funds for the purchase",
    };

    const errorMsg = errorStatusMessages[error.status] || "Error occurred during purchase";
    this.errorMessageToPrint = errorMsg;
  }

  onSellClick(){
    this.showSellInput = true;
    this.showBuyInput = false;
  }
  async Sell(quantity: number){ //generalize to all tickers
    //makes sure user is logged in beofre selling
    this.loginSignUpService.claimData().then((response) => {
      
    }).catch((error) => {
      console.log(error);
      if(error){
        this.errorMessageToPrint = "Please login first!"; //in the future this can be made a link
      }
    })
    const userData = await this.loginSignUpService.claimData();
    this.username = userData.username;
    this.shareAction.Sell(this.username, this.ticker, quantity).then((response) => { //make http request and wait for response, upon response send message to user
      if (response.status == 200) {
        this.errorMessageToPrint = "Shares successfully sold! Updating funds...";
      }

      setTimeout(() => {
            this.router.navigate(['.'], { relativeTo: this.route }).then(() => {
              location.reload();
            });
          }, 2000); // Delay of 2 seconds to read text
          /*
    Http status meanings in this function:
    404 - Username not found
    405 - Ticker not found
    406 - User doesn't any shares of the stock he wants to sell
    407 - Invalid shares quantity input
    408 - User trying to sell more shares than he owns
    */


    }).catch((error) => {
      console.log(error);
      this.handleSellError(error);
    });
    this.showSellInput = false; 
  }

  private handleSellError(error: any): void {
    const errorStatusMessages: Record<number, string> = {
      404: "Username not found",
      405: "Ticker not found",
      406: "You do not own any stocks to sell!",
      407: "Invalid quantity input!",
      408: "You are attempting to sell more shares than you own!",
    };

    const errorMsg = errorStatusMessages[error.status] || "Error occurred during purchase";
    this.errorMessageToPrint = errorMsg;
  }

}
