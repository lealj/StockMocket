import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class StocksService {
  private Url = `querystocks`;

  constructor(private http: HttpClient) { }

  getStockData(startMonth: number, startDay: number, startYear: number, endMonth: number, endDay: number, endYear: number): Observable<any[]> {
    const query = {
      ticker: 'MSFT',
      start_month: startMonth,
      start_day: startDay,
      start_year: startYear,
      end_month: endMonth,
      end_day: endDay,
      end_year: endYear
    };
    return this.http.post<any[]>(this.Url, query);
  }
}