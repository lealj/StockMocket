package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"gorm.io/gorm"
)

// For general stock information (fortune 500)

type Stock struct {
	gorm.Model
	Ticker string    `json:"ticker"`
	Price  float64   `json:"price"`
	Date   time.Time `json:"date"`
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

func GetStocks(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var stocks []Stock
	DB.Find(&stocks)
	json.NewEncoder(writer).Encode(stocks)
}

func GetStock(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	var stock []Stock

	DB.Where("ticker = ?", params["ticker"]).First(&stock)
	json.NewEncoder(writer).Encode(stock)
}

func UpdateStocks(writer http.ResponseWriter, rout *http.Request) {
	/*
		UpdateStockHelper("KO")
		UpdateStockHelper("MSFT")
		UpdateStockHelper("LMT")
		UpdateStockHelper("AAPL")
		UpdateStockHelper("WSF")
	*/
}

func UpdateStockHelper(tick string) {
	/*
		q, err := quote.Get(tick)
		if err != nil {
			fmt.Printf("Error getting quote: %v", err)
			return
		}

		var stockToUpdate Stock
		today := time.Now().Format("01-02-2006")

		DB.Model(&Stock{}).UpdateColumns(map[string]interface{}{
			"ticker": tick,
			"date":   today,
			"price":  q.RegularMarketPrice,
		})
	*/
}

func QueryStocks(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var query Query
	json.NewDecoder(router.Body).Decode(&query)

	// START AND END MUST NOT BE SOLEY ON A WEEKEND/HOLIDAY. Use same start/end values for a single day
	// Collects the dates and prices for the given start and end dates.
	p := &chart.Params{
		Symbol:   query.Ticker,
		Start:    &datetime.Datetime{Month: query.StartMonth, Day: query.StartDay, Year: query.StartYear},
		End:      &datetime.Datetime{Month: query.EndMonth, Day: query.EndDay, Year: query.EndYear},
		Interval: datetime.OneDay,
	}

	iter := chart.Get(p)
	// check error
	if iter.Err() != nil {
		writer.WriteHeader(http.StatusBadGateway)
		fmt.Printf("%v\n", iter.Err())
		return
	}

	// struct to store results
	var results []struct {
		Date  string  `json:"date"`
		Price float64 `json:"price"`
	}
	// Iterates over period start/end. Gets price and date through unix timestamp conversion
	for iter.Next() {
		close_price_f, _ := iter.Bar().Close.Float64()

		// time
		unix_timestamp := iter.Bar().Timestamp
		date := time.Unix(int64(unix_timestamp), 0).Format("01-02-2006")

		fmt.Printf("Price: %v Date: %v\n", close_price_f, date)

		// for return json data
		results = append(results, struct {
			Date  string  `json:"date"`
			Price float64 `json:"price"`
		}{Date: date, Price: close_price_f})
	}
	json.NewEncoder(writer).Encode(results)
}
