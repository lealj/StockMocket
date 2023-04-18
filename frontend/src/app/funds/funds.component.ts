import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'
import { Component, OnInit } from '@angular/core';
import { FundsService } from '../funds.service';

@Component({
  selector: 'app-funds',
  templateUrl: './funds.component.html',
  styleUrls: ['./funds.component.scss']
})
export class FundsComponent implements OnInit {

  showFunds: boolean = false;
  fundsText: number = 0;

  constructor (
    private loginSignUpService: LoginSignUpService,
    private fundsService : FundsService
  ) {}

  ngOnInit(): void {
    this.showFunds = false;
    this.getUserFunds();
  }

  async getUserFunds() { //it will be better to make a service to deal with updating the funds whenever buy/sell/logout happens. funds onyl updates when page is refreshed
    const response = await this.fundsService.getFunds();
    console.log(response);
    this.showFunds = true;
    this.fundsText = response;
  }

  
}
