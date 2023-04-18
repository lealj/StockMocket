package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// For keeping track of stocks a user owns.
type UserStocks struct {
	gorm.Model
	Username string `json:"username"`
	Ticker   string `json:"ticker"`
	Shares   int    `json:"shares"`
	// can add variable here summing the prices paid for the stocks, for the calculation of "gains/losses"
	// so 1 share bought at $50, another share bought at $75 -> $125 total. Future share worth $100. So profit = 100x2 - (50+75) = 200-125=$75
}

/*
Http status meanings in this function:
400 - Username not found
401 - Ticker not found
402 - Share quantity is not in range 1-50
403 - Not enough funds for the purchase
*/
func PurchaseStock(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Get username (from /userstock/{username} in post request header)

	// Get order info from post and assign username (we could just send the username in the body of the post request)
	var newPurchaseOrder UserStocks
	json.NewDecoder(router.Body).Decode(&newPurchaseOrder)

	fmt.Printf("Username: %s", newPurchaseOrder.Username)

	// Get credentials using username (may want to alter so that we don't get password here (bad practice potentially))
	var credentials Credentials
	if err := DB.Where("username = ?", newPurchaseOrder.Username).First(&credentials).Error; err != nil {
		fmt.Printf("Error finding username\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get funds
	funds := credentials.Funds

	fmt.Printf("Funds for user %s: $%f\n", newPurchaseOrder.Username, funds)

	// Get the stock data from Stocks (using ticker)
	var stock Stock
	if err := DB.Where("ticker = ?", newPurchaseOrder.Ticker).First(&stock).Error; err != nil {
		fmt.Printf("Error finding ticker\n")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Get price of stock and quantity wanted
	pricePerShare := stock.Price
	sharesInOrder := newPurchaseOrder.Shares

	// Base restriction on share purchase amount
	if sharesInOrder > 50 || sharesInOrder <= 0 {
		fmt.Printf("Invalid input for shares\n")
		writer.WriteHeader(http.StatusPaymentRequired)
		return
	}

	// Calculate total order cost
	totalOrderCost := float64(pricePerShare * float64(sharesInOrder))

	// Deny order if user doesn't have the funds
	if totalOrderCost > funds {
		fmt.Printf("Not enough funds for this order!\n")
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	// if reach this point, the purchase order is approved - create log
	CreateLog("buy", &newPurchaseOrder, totalOrderCost)

	var stocksOwned UserStocks
	// Check if user already owns shares of the company
	if err := DB.Where("username = ? AND ticker = ?", newPurchaseOrder.Username, newPurchaseOrder.Ticker).First(&stocksOwned).Error; err != nil {
		// Doesn't exist - create new purchase order entry in database
		fmt.Printf("User does not own %s already. Creating entry.\n", newPurchaseOrder.Ticker)
		DB.Create(&newPurchaseOrder)
	} else {
		// Does exist - update entry
		fmt.Printf("User owns %s already. Updating entry.\n", newPurchaseOrder.Ticker)
		stocksOwned.Shares = newPurchaseOrder.Shares + stocksOwned.Shares
		DB.Save(&stocksOwned)
	}

	// Update funds
	credentials.Funds = credentials.Funds - totalOrderCost
	// added this time to prevent error when test file used
	t := time.Now()
	t_f := t.Format("2006-01-02 15:04:05")
	createdAt, err := time.Parse("2006-01-02 15:04:05", t_f)
	if err != nil {
		fmt.Printf("error parsing time: %v", err)
	}
	credentials.CreatedAt = createdAt
	DB.Save(&credentials)

	var ret UserStocks
	DB.Where("username = ?", newPurchaseOrder.Username).First(&ret)

	json.NewEncoder(writer).Encode(ret)
}

/*
Http status meanings in this function:
404 - Username not found
405 - Ticker not found
406 - User doesn't any shares of the stock he wants to sell
407 - Invalid shares quantity input
408 - User trying to sell more shares than he owns
*/
func SellStock(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Get json data for struct
	var newSellOrder UserStocks
	json.NewDecoder(router.Body).Decode(&newSellOrder)

	// Get credentials using username (may want to alter so that we don't get password here (bad practice potentially))
	var credentials Credentials
	if err := DB.Where("username = ?", newSellOrder.Username).First(&credentials).Error; err != nil {
		fmt.Printf("Error finding username\n")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Get funds
	funds := credentials.Funds

	fmt.Printf("Current funds for user %s: $%f\n", newSellOrder.Username, funds)

	// Get the stock data from Stocks (using ticker)
	var stock Stock
	if err := DB.Where("ticker = ?", newSellOrder.Ticker).First(&stock).Error; err != nil {
		fmt.Printf("Error finding ticker\n")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Check that user actually has the shares he wants to sell
	var stocksOwned UserStocks
	// Check if user already owns shares of the company
	if err := DB.Where("username = ? AND ticker = ?", newSellOrder.Username, newSellOrder.Ticker).First(&stocksOwned).Error; err != nil {
		// Doesn't exist
		fmt.Printf("User does not own %s. Cannot sell what you don't own!\n", newSellOrder.Ticker)
		writer.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// Base restriction on share purchase amount
	if newSellOrder.Shares > 50 || newSellOrder.Shares <= 0 {
		fmt.Printf("Invalid input for shares\n")
		writer.WriteHeader(http.StatusProxyAuthRequired)
		return
	}

	// Does exist
	fmt.Printf("User owns %d shares of %s.\n", stocksOwned.Shares, newSellOrder.Ticker)
	// Check that shares selling is <= shares owned
	if stocksOwned.Shares < newSellOrder.Shares {
		// User selling more shares than he owns
		fmt.Printf("%s attempting to sell more shares of %s than he owns!\n", newSellOrder.Username, newSellOrder.Ticker)
		writer.WriteHeader(http.StatusRequestTimeout)
		return
	}
	// Update Funds
	pricePerShare := stock.Price
	sharesInOrder := newSellOrder.Shares
	totalOrderValue := float64(pricePerShare * float64(sharesInOrder))
	credentials.Funds = credentials.Funds + totalOrderValue
	DB.Save(&credentials)

	// create log
	CreateLog("sell", &newSellOrder, totalOrderValue)

	// Update share count
	if stocksOwned.Shares-newSellOrder.Shares == 0 {
		// Delete entry in database if 0
		DB.Unscoped().Delete(&stocksOwned)
		fmt.Printf("Sell complete. User now owns 0 shares of %s. User funds is now: $%f\n\n",
			stocksOwned.Ticker, credentials.Funds)
	} else {
		stocksOwned.Shares = stocksOwned.Shares - newSellOrder.Shares
		DB.Save(&stocksOwned)
		fmt.Printf("Sell complete. User now owns %d shares of %s. User funds is now: $%f\n\n",
			stocksOwned.Shares, stocksOwned.Ticker, credentials.Funds)
	}

	json.NewEncoder(writer).Encode(stocksOwned)
}

type StockTickerShares struct {
	Ticker string  `json:"ticker"`
	Shares int     `json:"shares"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"` //a percent
}

// this returns slices of information in the format of the above struct
func GetStocksOwned(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// username
	var userstock UserStocks
	json.NewDecoder(router.Body).Decode(&userstock)

	user_stocks := GetUserStocksArray(userstock.Username)

	// create empty slice and populate with info from user_stocks
	stocks := make([]StockTickerShares, 0)
	for _, s := range user_stocks {
		price := GetStockPrice(s.Ticker)
		change := GetIndividualStockChange(s.Username, s.Ticker, s.Shares)
		stock := StockTickerShares{Ticker: s.Ticker, Shares: s.Shares, Price: price, Change: change}
		stocks = append(stocks, stock)
	}

	json.NewEncoder(writer).Encode(stocks)
}

func GetUserStocksArray(username string) []UserStocks {
	var user_stocks []UserStocks
	DB.Where("username = ?", username).Find(&user_stocks)
	return user_stocks
}

func ResetAccount(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// get credentials
	var creds Credentials
	json.NewDecoder(router.Body).Decode(&creds)
	//log.Printf("Username to reset: %s", creds.Username)

	// reset funds to default value
	err2 := DB.Where("username = ?", creds.Username).First(&creds).Error
	if err2 != nil {
		fmt.Printf("Username not found in credentials: %v", err2)
		http.Error(writer, "Username not found in credentials", http.StatusBadRequest)
		return
	}
	creds.Funds = 1000
	DB.Save(&creds)

	deletion := DB.Where("username = ?", creds.Username).Unscoped().Delete(&UserStocks{})
	if deletion.Error != nil {
		fmt.Printf("Error finding username during deletion (userstocks): %v\n", deletion.Error)
		http.Error(writer, "Error during deletion (userstocks)", http.StatusBadRequest)
		return
	}

	DeleteAllLogs(creds.Username)

	writer.WriteHeader(http.StatusOK)

}
