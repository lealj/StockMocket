package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

// For general stock information (fortune 500)

type Stock struct {
	gorm.Model
	Ticker string  `json:"ticker"`
	Price  float32 `json:"price"`
}

type Stock_no_gorm struct {
	Ticker string  `json:"ticker"`
	Price  float32 `json:"price"`
}

type Query struct {
	Ticker     string `json:"ticker"`
	StartMonth int    `json:"start_month"`
	StartDay   int    `json:"start_day"`
	StartYear  int    `json:"start_year"`
	EndMonth   int    `json:"end_month"`
	EndDay     int    `json:"end_day"`
	EndYear    int    `json:"end_year"`
}

func SaveStockPrice(s Stock) {
	// load entry for stock in db
	var stock_db_entry Stock
	DB.Where("ticker = ?", s.Ticker).First(&stock_db_entry)
	// update price
	stock_db_entry.Price = s.Price
	// update time
	t := time.Now()
	t_f := t.Format("2006-01-02 15:04:05")
	createdAt, err := time.Parse("2006-01-02 15:04:05", t_f)
	if err != nil {
		fmt.Printf("error parsing time: %v", err)
	}
	// save in database
	stock_db_entry.CreatedAt = createdAt
	DB.Save(&stock_db_entry)
}

func GetStockPrice(ticker string) float32 {
	var price float32
	DB.Table("stocks").Where("ticker = ?", ticker).Select("price").Scan(&price)
	return price
}

func GetStocks(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var stocks []Stock
	var stocksResponse []Stock_no_gorm
	DB.Find(&stocks).Select("ticker, price").Scan(&stocksResponse)
	json.NewEncoder(writer).Encode(stocksResponse)
}

func GetStock(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var stock Stock
	json.NewDecoder(rout.Body).Decode(&stock)
	DB.Where("ticker = ?", stock.Ticker).Find(&stock)
	json.NewEncoder(writer).Encode(stock)
}

func InitFinnhubClient() *finnhub.DefaultApiService {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", "ci87gu9r01qnrgm32fggci87gu9r01qnrgm32fh0")
	return finnhub.NewAPIClient(cfg).DefaultApi
}

func UpdateStocks() {
	// init client
	finnhubClient := InitFinnhubClient()

	tickers := []string{"KO", "MSFT", "LMT", "AAPL", "WFC"}
	for _, ticker := range tickers {
		res, _, err := finnhubClient.Quote(context.Background()).Symbol(ticker).Execute()
		if err != nil {
			fmt.Printf("Error Updating %s: %v\n", ticker, err)
			continue
		}

		var stock Stock
		DB.Where("ticker = ?", ticker).First(&stock)
		stock.Price = float32(*res.C)

		SaveStockPrice(stock)
	}
	fmt.Println("Stocks updated.")
}

type Data struct {
	Values []float32 `json:"values"`
	Dates  []string  `json:"dates"`
}

func QueryStocks(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var query Query
	json.NewDecoder(router.Body).Decode(&query)

	unix_start_interval := time.Date(query.StartYear, time.Month(query.StartMonth), query.StartDay, 0, 0, 0, 0, time.UTC).Unix()
	unix_end_interval := time.Date(query.EndYear, time.Month(query.EndMonth), query.EndDay, 0, 0, 0, 0, time.UTC).Unix()

	// init client
	finnhubClient := InitFinnhubClient()

	res, _, err := finnhubClient.StockCandles(context.Background()).Symbol(query.Ticker).Resolution("D").From(unix_start_interval).To(unix_end_interval).Execute()

	if err != nil {
		fmt.Printf("Error with stock candles: %v\n", err)
	} else {
		//get values from *res.C and prepare to send json data.
		unix_dates := *res.T
		var dates []string
		dateFormat := "01-02-2006"
		for _, value := range unix_dates {
			dateObj := time.Unix(value, 0)
			dateStr := dateObj.Format(dateFormat)
			dates = append(dates, dateStr)
		}
		data := Data{
			Values: *res.C,
			Dates:  dates,
		}
		json.NewEncoder(writer).Encode(data)
	}
}
