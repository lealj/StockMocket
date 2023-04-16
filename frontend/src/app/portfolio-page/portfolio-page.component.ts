import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { CookieServices } from "../cookie.service";
import {catchError} from "rxjs";
import { PortfolioPageService } from './portfolio-page.service'

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

  constructor(private portfolioPageService: PortfolioPageService) {}

  ngOnInit(): void {
    this.GetPortfolioValue()
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
          console.log("Display pv faileddf")
        } else {
          console.log("Display pv failed")
        }
      }
    );
  }
}
