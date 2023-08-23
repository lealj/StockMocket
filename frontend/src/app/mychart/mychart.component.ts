import { Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { StocksService } from '../services/stocks.service';
import { Message, WebsocketService } from '../services/websocket.service';
import { ActivatedRoute } from '@angular/router';

Chart.register(...registerables);

@Component({
  selector: 'app-mychart',
  templateUrl: './mychart.component.html',
  styleUrls: ['./mychart.component.scss'],
  providers: [WebsocketService],
})

export class MychartComponent implements OnInit {
    public stockData: any[] = [];
    public content: number = 0;
    received: any[]=[];
    public ticker: string = '';
    private myChart: Chart | null = null;

    constructor(
      private stocksService: StocksService, 
      public WebsocketService: WebsocketService,
      private route: ActivatedRoute
    ) 
    {
      this.route.paramMap.subscribe(params => {
        const tickerParam = params.get('ticker');
        this.ticker = tickerParam !== null ? tickerParam: '';
      });
      WebsocketService.messages.subscribe(msg => {
        if (msg.source == this.ticker) {
          this.content = parseFloat(msg.content);
          this.updateLastPrice(this.content);
        }
      });
    }

    ngOnInit(): void {
      const startDate = new Date();
      startDate.setMonth(startDate.getMonth() - 6);
      this.stocksService.getStockData(this.ticker, startDate.getMonth() + 1, startDate.getDate(), startDate.getFullYear(), new Date().getMonth() + 1, new Date().getDate(), new Date().getFullYear())
        .subscribe(data => {
          this.stockData = data;
          // break down stockData into values and dates
          const { values, dates } = this.stockData as unknown as { values: number[], dates: any[] };
          this.content = values[values.length -1];

          // Extract the month and year from each date and use them as the label
          const monthLabels = dates.map(date => {
            const [month, day, year] = date.split('-');
            const monthName = new Date(parseInt(year), parseInt(month) - 1, parseInt(day)).toLocaleString('en-US', { month: 'short'});
            return `${monthName} ${year}`;
          });
          
          // Extract the month and year from each date and use as y labels
          this.myChart = new Chart('myChart', {
            type: 'line',
            data: {
              labels: monthLabels,
              datasets: [
                {
                  label: this.ticker,
                  data: values,
                  pointRadius: 0.01,
                  backgroundColor: 'rgba(0, 255, 0, 1)',
                  borderColor: 'rgba(0, 255, 0, 1)',
                  borderWidth: 2,
                }
              ]
            },
            options: {
              plugins: {
                legend: {
                  display: false
                },
                tooltip: {
                  intersect: false,
                  mode: 'index',
                  position: 'nearest',
                  displayColors: false,
                  callbacks: {
                    label: (tooltipItem: any) => {
                      if (tooltipItem.dataset.label && tooltipItem.parsed) {
                        const date = new Date(dates[tooltipItem.parsed.x]);
                        const formattedDate = date.toLocaleDateString('en-US', {
                          month: 'short',
                          day: 'numeric',
                          year: 'numeric'
                        });
                        const value = tooltipItem.formattedValue || '';
                        return `$${value} ${formattedDate}`;
                      }
                      return '';
                    }
                  }
                }
              },
              scales: {
                x: {
                  ticks: {
                    maxTicksLimit: 6
                  },
                },
                y: {
                  beginAtZero: false,
                }
              }
            }
          });
        });
    }

    updateLastPrice(newPrice: number) {
      if(this.myChart !== null && this.myChart.data !== null) {
        const lastDataset = this.myChart.data.datasets[0];
        const lastDataIndex = lastDataset.data.length -1;

        lastDataset.data[lastDataIndex] = newPrice;

        this.myChart.update();
        this.myChart.render();
      }
    }
}


