import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import {JwtModule} from '@auth0/angular-jwt';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MychartComponent } from './mychart/mychart.component';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';
import { HeaderComponent } from './header/header.component';
import { BuySellButtonComponent } from './buy-sell-button/buy-sell-button.component';
import { HomeComponent } from './home/home.component';
import { PortfolioPageComponent } from './portfolio-page/portfolio-page.component';
import { AboutPageComponent } from './about-page/about-page.component';
import { LogOutPageComponent } from './log-out-page/log-out-page.component';
import { FundsComponent } from './funds/funds.component';

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
    PortfolioPageComponent,
    AboutPageComponent,
    LogOutPageComponent,
    FundsComponent,
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
