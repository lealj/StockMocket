import { Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { StocksService } from '../stocks.service';

Chart.register(...registerables);

@Component({
  selector: 'app-mychart',
  templateUrl: './mychart.component.html',
  styleUrls: ['./mychart.component.scss']
})
export class MychartComponent implements OnInit {

    public stockData: any[] = [];

    constructor(private stocksService: StocksService) { }

    ngOnInit(): void {
        const startDate = new Date();
        startDate.setMonth(startDate.getMonth() - 4);
        this.stocksService.getStockData(startDate.getMonth() + 1, startDate.getDate(), startDate.getFullYear(), new Date().getMonth() + 1, new Date().getDate(), new Date().getFullYear())
          .subscribe(data => {
            this.stockData = data;
            const dates = this.stockData.map(item => item.date);
            const prices = this.stockData.map(item => item.price);
      
            const myChart = new Chart('myChart', {
              type: 'line',
              data: {
                labels: dates,
                datasets: [
                  {
                    label: 'MSFT',
                    data: prices,
                    pointRadius: 2, 
                    backgroundColor: [
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)'
                    ],
                    borderColor: [
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)',
                      'rgba(0, 255, 0, 1)'
                    ],
                    borderWidth: 1
                  }
                ]
              },
              options: {
                scales: {
                  y: {
                    beginAtZero: false,
                    min: 200
                  }
                }
              }
            });
          });
    }
    
}


