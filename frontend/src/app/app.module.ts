import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { JwtModule } from '@auth0/angular-jwt';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MychartComponent } from './components/mychart/mychart.component';
import { LoginsignuppageComponent } from './components/login-signup-page/loginsignuppage.component';
import { HeaderComponent } from './components/header/header.component';
import { BuySellButtonComponent } from './buy-sell-button/buy-sell-button.component';
import { LandingComponent } from './components/landing-page/landing.component';
import { AboutPageComponent } from './components/about-page/about-page.component';
import { LogOutPageComponent } from './components/log-out-page/log-out-page.component';
import { FundsComponent } from './components/funds/funds.component';
import { PortfolioPageComponent } from './components/portfolio-page/portfolio-page.component';
import { StocksPageComponent } from './components/stocks-page/stocks-page.component';

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
    LandingComponent,
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
