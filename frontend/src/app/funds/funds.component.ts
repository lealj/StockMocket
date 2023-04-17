import { LoginSignUpService } from '../loginsignuppage/loginsignuppage.service'
import { Component, OnInit,  ChangeDetectorRef } from '@angular/core';

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
    private changeDetectorRef: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    this.showFunds = false;
      this.getUserFunds();
  }

  async getUserFunds() { //it will be better to make a service to deal with updating the funds whenever buy/sell/logout happens. funds onyl updates when page is refreshed
    const response = await this.loginSignUpService.getFunds();
    console.log(response);
    this.showFunds = true;
    this.fundsText = response;
    this.changeDetectorRef.detectChanges();
  }
  
}
