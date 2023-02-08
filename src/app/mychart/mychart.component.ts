import { Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
Chart.register(...registerables);

@Component({
  selector: 'app-mychart',
  templateUrl: './mychart.component.html',
  styleUrls: ['./mychart.component.scss']
})
export class MychartComponent implements OnInit {
  ngOnInit(){
    var myChart = new Chart("myChart", {
      type: 'line',
      data: {
          labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
          datasets: [{
              label: 'MSFT',
              data: [12, 19, 3, 5, 2, 3],
              backgroundColor: [
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)'
              ],
              borderColor: [
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)',
                  'rgba(255, 0, 0, 1)'
              ],
              borderWidth: 1
          }]
      },
      options: {
          scales: {
              y: {
                  beginAtZero: true
              }
          }
      }
  });
  }
  
}


