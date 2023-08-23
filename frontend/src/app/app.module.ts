import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { JwtModule } from '@auth0/angular-jwt';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MychartComponent } from './mychart/mychart.component';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';
import { HeaderComponent } from './header/header.component';
import { BuySellButtonComponent } from './buy-sell-button/buy-sell-button.component';
import { HomeComponent } from './home/home.component';
import { AboutPageComponent } from './about-page/about-page.component';
import { LogOutPageComponent } from './log-out-page/log-out-page.component';
import { FundsComponent } from './funds/funds.component';
import { PortfolioPageComponent } from './portfolio-page/portfolio-page.component';
import { StocksPageComponent } from './stocks-page/stocks-page.component';

export function tokenGetter() {
  return localStorage.getItem('token');
}

@NgModule({
  declarations: [
    AppComponent,
    MychartComponent,
    LoginsignuppageComponent,
    HeaderComponent,
    BuySellButtonComponent,
    HomeComponent,
    AboutPageComponent,
    PortfolioPageComponent,
    LogOutPageComponent,
    FundsComponent,
    StocksPageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: tokenGetter
      },
    }),
  ],
  exports: [
    BuySellButtonComponent,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
