package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// For keeping track of stocks a user owns.
type UserStocks struct {
	gorm.Model
	Username	string `json:"username"`
	Ticker 		string `json:"ticker"`
	Shares 		int   `json:"shares"`
	// can add variable here summing the prices paid for the stocks, for the calculation of "gains/losses"
	// so 1 share bought at $50, another share bought at $75 -> $125 total. Future share worth $100. So profit = 100x2 - (50+75) = 200-125=$75
}

/* To keep a log of purchases (for future use potentially)
type PurchaseHistory struct {
	gorm.Model
	Username	string `json:"username"`
	Ticker 		string `json:"ticker"`
	Shares 		int   `json:"shares`
}
*/

func PurchaseStock(writer http.ResponseWriter, router *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	
	// Get username (from /userstock/{username} in post request header)
	params := mux.Vars(router)
	username := params["username"]

	// Get order info from post and assign username (we could just send the username in the body of the post request)
	var newPurchaseOrder UserStocks
	json.NewDecoder(router.Body).Decode(&newPurchaseOrder)
	newPurchaseOrder.Username = username

	// Get credentials using username (may want to alter so that we don't get password here (bad practice potentially))
	var credentials Credentials
    if err := DB.Where("username = ?", newPurchaseOrder.Username).First(&credentials).Error; err != nil {
		fmt.Printf("Error finding username\n")
		writer.WriteHeader(http.StatusBadRequest)
        return
    }
	// Get funds
	funds := credentials.Funds

	fmt.Printf("Funds for user %s: $%d\n", newPurchaseOrder.Username, funds)

	// Get the stock data from Stocks (using ticker)
	var stock Stock
	if err := DB.Where("ticker = ?", newPurchaseOrder.Ticker).First(&stock).Error; err != nil {
		fmt.Printf("Error finding ticker\n")
		writer.WriteHeader(http.StatusBadRequest)
        return
    }
	// Get price of stock and quantity wanted
	pricePerShare := stock.LatestPrice
	sharesInOrder := newPurchaseOrder.Shares

	// Calculate total order cost
	totalOrderCost := pricePerShare * sharesInOrder
	
	// Deny order if user doesn't have the funds 
	if totalOrderCost > funds {
		fmt.Printf("Not enough funds for this order!\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	} 

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
	fmt.Printf("New funds for %s: $%d\n", username, credentials.Funds)
	DB.Save(&credentials)
}

func SellStock(writer http.ResponseWriter, router *http.Request){
	writer.Header().Set("Content-Type", "application/json")

	// Get username
	params := mux.Vars(router)
	username := params["username"]

	// Get json data for struct
	var newSellOrder UserStocks
	json.NewDecoder(router.Body).Decode(&newSellOrder)
	newSellOrder.Username = username

	// Get credentials using username (may want to alter so that we don't get password here (bad practice potentially))
	var credentials Credentials
    if err := DB.Where("username = ?", newSellOrder.Username).First(&credentials).Error; err != nil {
		fmt.Printf("Error finding username\n")
		writer.WriteHeader(http.StatusBadRequest)
        return
    }

	// Get funds
	funds := credentials.Funds

	fmt.Printf("Current funds for user %s: $%d\n", newSellOrder.Username, funds)

	// Get the stock data from Stocks (using ticker)
	var stock Stock
	if err := DB.Where("ticker = ?", newSellOrder.Ticker).First(&stock).Error; err != nil {
		fmt.Printf("Error finding ticker\n")
		writer.WriteHeader(http.StatusBadRequest)
        return
    }

	// Check that user actually has the shares he wants to sell
	var stocksOwned UserStocks
	// Check if user already owns shares of the company
	if err := DB.Where("username = ? AND ticker = ?", newSellOrder.Username, newSellOrder.Ticker).First(&stocksOwned).Error; err != nil {
		// Doesn't exist
		fmt.Printf("User does not own %s. Cannot sell what you don't own!\n", newSellOrder.Ticker)
		writer.WriteHeader(http.StatusBadRequest)
		return 
	} 
	// Does exist
	fmt.Printf("User owns %d shares of %s.\n", stocksOwned.Shares, newSellOrder.Ticker)
	// Check that shares selling is <= shares owned
	if stocksOwned.Shares < newSellOrder.Shares {
		// User selling more shares than he owns
		fmt.Printf("%s attempting to sell more shares of %s than he owns!\n", newSellOrder.Username, newSellOrder.Ticker)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Update Funds
	pricePerShare := stock.LatestPrice
	sharesInOrder := newSellOrder.Shares
	totalOrderValue := pricePerShare * sharesInOrder
	credentials.Funds = credentials.Funds + totalOrderValue
	DB.Save(&credentials)

	// Update share count
	if stocksOwned.Shares - newSellOrder.Shares == 0 {
		// Delete entry in database if 0
		DB.Unscoped().Delete(&stocksOwned)
		fmt.Printf("Sell complete. User now owns 0 shares of %s. User funds is now: $%d\n\n",
			stocksOwned.Shares, stocksOwned.Ticker, credentials.Funds)
	} else {
		stocksOwned.Shares = stocksOwned.Shares - newSellOrder.Shares
		DB.Save(&stocksOwned)
		fmt.Printf("Sell complete. User now owns %d shares of %s. User funds is now: $%d\n\n", 
			stocksOwned.Shares, stocksOwned.Ticker, credentials.Funds)
	}
	
}

func GetStocksOwned(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	var user_stocks []UserStocks
	DB.Where("username = ?", params["username"]).Find(&user_stocks)
	// this can be returned in a better format, or it can be parsed in front end. 
	json.NewEncoder(writer).Encode(user_stocks)
}
