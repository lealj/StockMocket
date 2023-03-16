package main

import (
	"gorm.io/gorm"
)

// For keeping track of stocks a user owns.

type UserStocks struct {
	gorm.Model
	Username        string `json:"username"`
	TickerAndShares []UserStockHelper
}

type UserStockHelper struct {
	gorm.Model
	Ticker      string
	ShareAmount uint
}

// func PurchaseStock(writer http.ResponseWriter, router *http.Request){}

// func SellStock(writer http.ResponseWriter, router *http.Request){}

// func GetStocksOwned(writer http.ResponseWriter, router *http.Request) {}
