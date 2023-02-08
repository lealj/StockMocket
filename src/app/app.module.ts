import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MychartComponent } from './mychart/mychart.component';
import { LoginsignuppageComponent } from './loginsignuppage/loginsignuppage.component';

@NgModule({
  declarations: [
    AppComponent,
    MychartComponent,
    LoginsignuppageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
