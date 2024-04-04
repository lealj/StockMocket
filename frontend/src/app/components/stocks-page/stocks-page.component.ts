import { Component, OnInit } from '@angular/core';
import { StocksPageService } from '../../services/page-services/stocks-page.service'

interface Stock
{
    ticker: string;
    price: number;
}

@Component({
    selector: 'app-stocks-page',
    templateUrl: './stocks-page.component.html',
    styleUrls: ['./stocks-page.component.scss'],
    providers: [StocksPageService]
})

export class StocksPageComponent implements OnInit{
    // variables
    public stocksList: Stock[] = [] 
    //construct
    constructor(private stocksPageService: StocksPageService) {}

    ngOnInit(): void {
        this.GetStocksList() 
    }
    // functions
    GetStocksList() {
        this.stocksPageService.getStocksList().then((response) => {
            this.stocksList = response.body;
        })
        .catch((error) => {

        });
    }
}