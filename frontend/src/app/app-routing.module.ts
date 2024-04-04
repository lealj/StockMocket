import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginsignuppageComponent } from './components/login-signup-page/loginsignuppage.component';
import { MychartComponent } from './components/mychart/mychart.component';
import { LandingComponent } from './components/landing-page/landing.component';
import { AboutPageComponent } from './components/about-page/about-page.component';
import { LogOutPageComponent } from './components/log-out-page/log-out-page.component';
import { PortfolioPageComponent } from './components/portfolio-page/portfolio-page.component'
import { StocksPageComponent } from './components/stocks-page/stocks-page.component';
 
const routes: Routes = [
  {
    path: '',
    component: LandingComponent
  },
  {
    path: 'account',
    component: LoginsignuppageComponent
  },
  {
    path: 'portfolio',
    component: PortfolioPageComponent
  },
  {
    path: 'about',
    component: AboutPageComponent
  },
  {
    path: 'logout',
    component: LogOutPageComponent
  },
  {
    path: 'viewstocks',
    component: StocksPageComponent
  },
  {
    path: 'viewstocks/:ticker', 
    component: MychartComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
