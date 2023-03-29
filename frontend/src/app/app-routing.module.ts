import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';
import { MychartComponent } from './mychart/mychart.component';

const routes: Routes = [
  {
    path: '',
    component: LoginsignuppageComponent
  },
  {
    path: 'charts',
    component: MychartComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
