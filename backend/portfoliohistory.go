package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// To keep track of history - a log book essentially
type PortfolioHistory struct {
	gorm.Model
	Username  string  `json:"username"`
	Ticker    string  `json:"ticker"`
	Shares    int     `json:"shares"`
	OrderType string  `json:"ordertype"` // should be buy, sell, div
	Value     float64 `json:"price"`
}

func CreateLog(ordertype string, order *UserStocks, value float64) {
	var log PortfolioHistory
	log.Username = order.Username
	log.Ticker = order.Ticker
	log.Shares = order.Shares
	log.OrderType = ordertype
	log.Value = value
	DB.Create(&log)
}

func DeleteAllLogs(username string) {
	// delete logs associated with this username
	delete_logs := DB.Where("username = ?", username).Unscoped().Delete(&PortfolioHistory{})
	if delete_logs.Error != nil {
		fmt.Printf("Error during logs deletion: %v\n", delete_logs.Error)
	}
}

func GetLogs(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// get username
	var log PortfolioHistory
	json.NewDecoder(router.Body).Decode(&log)

	// find logs containing username
	var user_logs []PortfolioHistory
	DB.Where("username = ?", log.Username).Find(&user_logs)
	// this can be returned in a better format, or it can be parsed in front end.
	json.NewEncoder(writer).Encode(user_logs)
}

/* For Portfolio page */

type PortfolioInfo struct {
	PortfolioValue float64 `json:"portfolio_value"`
	PVChange       float64 `json:"pv_change"`
	// possibly add invidividual stock data pertaining to the user here (in a string)
	// example: avg cost for x stock, percent change etc. see getindividualstockchange function
}

func GetUserPortfolioInfo(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// get username
	var log PortfolioHistory
	json.NewDecoder(router.Body).Decode(&log)

	portfolio_value := GetUserPortfolioValue(log.Username)
	pv_change_percent := ((portfolio_value - 1000) / 1000) * 100
	// indiv stock info

	portfolio_info := PortfolioInfo{
		PortfolioValue: portfolio_value,
		PVChange:       pv_change_percent,
	}

	json.NewEncoder(writer).Encode(portfolio_info)
}

// Gets value of portfolio based on stock price * shares owned + funds
func GetUserPortfolioValue(username string) float64 {
	// get user's stocks
	user_stocks := GetUserStocksObject(username)

	// goes through stocks the user owns, multiplies their price by shares, sums.
	pv := 0.00
	for i := range user_stocks {
		var stock Stock
		DB.Where("ticker = ?", user_stocks[i].Ticker).First(&stock)
		pv += stock.Price * float64(user_stocks[i].Shares)
	}

	// gets funds from credentials based on username
	var funds float64
	DB.Table("credentials").Where("username = ?", username).Select("funds").Scan(&funds)
	pv += funds

	return pv
}

/*
func GatherStocksChangeData(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// use below function
}


// Gets the difference between the avg purchase price * shares and current purchase price * shares

func GetIndividualStockChange(user_stock UserStocks) float64 {
	// get logs corresponding to the user and the ticker.
	var logs []PortfolioHistory
	DB.Where("username = ?", user_stock.Username).Where("ticker = ?", user_stock.Ticker).Find(&logs)

	// get sum of purchases
	var total_cost float64
	var total_shares_bought int
	for i := range logs {
		if logs[i].OrderType == "buy" {
			total_cost += logs[i].Value
			total_shares_bought += logs[i].Shares
		}
	}
	// calculate avg cost per share
	avg_share_price := total_cost / float64(total_shares_bought)
	avg_cost_total := avg_share_price * float64(user_stock.Shares)

	// current value of the users stocks (current price * shares)
	var current_value float64
	DB.Table("stocks").Where("ticker = ?", user_stock.Ticker).Select("price").Scan(&current_value)
	current_value = current_value * float64(user_stock.Shares)

	// calculate change (new - old)/old
	return (current_value - avg_cost_total) / avg_cost_total
}
*/

func GetUserStocksObject(username string) []UserStocks {
	var user_stocks []UserStocks
	DB.Where("username = ?", username).Find(&user_stocks)
	return user_stocks
}
