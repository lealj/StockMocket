import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';
import { MychartComponent } from './mychart/mychart.component';
import { HomeComponent } from './home/home.component';
import { AboutPageComponent } from './about-page/about-page.component';
import { LogOutPageComponent } from './log-out-page/log-out-page.component';
import { PortfolioPageComponent } from './portfolio-page/portfolio-page.component'
import { StocksPageComponent } from './stocks-page/stocks-page.component';
 
const routes: Routes = [
  {
    path: '',
    component: HomeComponent
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
