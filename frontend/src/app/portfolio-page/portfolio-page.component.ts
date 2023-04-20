import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { CookieServices } from "../cookie.service";
import {catchError} from "rxjs";
import { PortfolioPageService } from './portfolio-page.service'

interface Stock 
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
  providers: [PortfolioPageService]
})

export class PortfolioPageComponent implements OnInit{
  public response: any;
  public portfolioValue: any;
  public portfolioChange: any;
  public stocksOwned: Stock[] = []
  public logs: Log[] = []

  constructor(private portfolioPageService: PortfolioPageService) {}

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
      console.log
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


