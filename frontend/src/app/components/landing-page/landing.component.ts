import { Component } from '@angular/core';
@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss']
})

export class LandingComponent {
  computerGraphic: string = '../../../assets/computer_graphic.svg';
  bookGraphic: string = '../../../assets/books_graphic.svg';
  analyticsGraphic: string = '../../../assets/analytics_graphic.svg';
  cabinetGraphic: string = '../../../assets/cabinet_graphic.svg';
}
