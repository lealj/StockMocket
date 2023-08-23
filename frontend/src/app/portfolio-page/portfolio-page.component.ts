import { Component, OnInit } from '@angular/core';
import { PortfolioPageService } from './portfolio-page.service'
import { WebsocketService } from '../services/websocket.service';

interface StockTransaction
{
  ticker: string;
  shares: number;
  price: number;
  change: number;
}

interface Log
{
  Date: Date;
  ticker: string;
  shares: number;
  ordertype: string;
  price: number;
}

@Component({
  selector: 'app-portfolio-page',
  templateUrl: './portfolio-page.component.html',
  styleUrls: ['./portfolio-page.component.scss'],
  providers: [PortfolioPageService, WebsocketService]
})

export class PortfolioPageComponent implements OnInit{
  public response: any;
  public portfolioValue: any;
  public portfolioChange: any;
  public stocksOwned: StockTransaction[] = []
  public logs: Log[] = []
  private ogPrice = 0;

  constructor(
    private portfolioPageService: PortfolioPageService,
    private WebsocketService: WebsocketService
  ) 
  {
    WebsocketService.messages.subscribe(msg => {
      let ticker: string = msg.source;
      let price: number = parseFloat(msg.content);
      const updatedStock = this.stocksOwned.find(stock => stock.ticker === ticker);
      if (updatedStock) {
        if(this.ogPrice === 0) {
          this.ogPrice = updatedStock.price / (1 + (updatedStock.change/100));
        }
        updatedStock.change = ((price - this.ogPrice)/this.ogPrice)*100
        updatedStock.price = price;
      }
    })
  }

  ngOnInit(): void {
    this.GetPortfolioValue()
    this.GetOwnedStocks()
    this.GetUserLogs()
  }

  GetOwnedStocks() {
    this.portfolioPageService.getOwnedStocks().then((response) => {
      this.stocksOwned = response.body; 
    })
    .catch((error) => {
        if(error.error === null)
        {
          console.log("null")
        } else {
          console.log("get owned stocks failed")
        }
      }
    );
  }

  GetUserLogs()
  {
    this.portfolioPageService.getUserLogs().then((response) => {
      this.logs = response.body;
    })
    .catch((error) => {
        if(error.error === null)
        {
          console.log("null")
        } else {
          console.log("get user logs failed")
        }
      }
    );
  }

  GetPortfolioValue()
  {
    this.portfolioPageService.getPortfolioValue().then((response) => {
      this.portfolioValue = response.body.portfolio_value;
      this.portfolioChange = response.body.pv_change;
    })
    .catch((error) => {
        if(error.error === null)
        {
          console.log("null")
        } else {
          console.log("Display pv failed")
        }
      }
    );
  }
}