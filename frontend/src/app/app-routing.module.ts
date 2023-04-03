import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';
import { MychartComponent } from './mychart/mychart.component';
import { HomeComponent } from './home/home.component';
import { PortfolioPageComponent } from './portfolio-page/portfolio-page.component';
import { AboutPageComponent } from './about-page/about-page.component';

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
    path: 'charts',
    component: MychartComponent
  },
  {
    path: 'portfolio',
    component: PortfolioPageComponent
  },
  {
    path: 'about',
    component: AboutPageComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
