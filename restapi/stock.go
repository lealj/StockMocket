package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	CompanyName string `json:"companyname"`
	Ticker      string `json:"ticker"`
	LatestPrice string `json:"price"`
	Date        string `json:"date"`
}

/*
// schema for individual stock (???)
type Stock struct {
	gorm.Model
	Ticker string `json:"ticker"`
	Price  string `json:"price"`
	Date   string `json:"date"`
}
*/

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
	DB.First(&stock, params["tick"])
	json.NewEncoder(writer).Encode(stock)
}
