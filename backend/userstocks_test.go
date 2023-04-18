package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/gorm"
)

func TestPurchaseStock(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testPurchase := `{"username":"john", "ticker":"MSFT", "shares":1}`

	var cred Credentials
	DB.Where("username = ?", "john").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	r := httptest.NewRequest("POST", "/userstock/buy", strings.NewReader(testPurchase))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	PurchaseStock(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains correct parameters. see "expect"
	expect := `"username":"john","ticker":"MSFT","shares":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

// Attempts purchase without necessary funds
func TestPurchaseStock_NoFunds(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testPurchase := `{"username":"nofunds", "ticker":"MSFT", "shares":3}`

	var cred Credentials
	DB.Where("username = ?", "nofunds").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	r := httptest.NewRequest("POST", "/userstock/buy", strings.NewReader(testPurchase))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	PurchaseStock(w, r)

	// check the response - should be 400, so err if statusok (200)
	if status := w.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	// check that there is no ownership entry in user_stocks
	var nofunds_ownership UserStocks

	if err := DB.Where("username = ?", "nofunds").First(&nofunds_ownership).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Username not found")
		}
	} else {
		t.Errorf("Record was found in user_stocks. This should not be the case.")
	}
}

func TestSellStock(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testSale := `{"username":"john", "ticker":"MSFT", "shares":1}`

	// Get credentials for username and funds
	var cred Credentials
	DB.Where("username = ?", "john").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	// set reader and writer
	r := httptest.NewRequest("POST", "/userstock/sell", strings.NewReader(testSale))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	SellStock(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains correct parameters. see "expect"
	expect := `"username":"john","ticker":"MSFT","shares":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

// Tests selling stock when user doesn't own it
func TestSellStock_NotOwned(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testSale := `{"username":"nofunds", "ticker":"MSFT", "shares":1}`

	var cred Credentials
	DB.Where("username = ?", "nofunds").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	// set reader and writer
	r := httptest.NewRequest("POST", "/userstock/sell", strings.NewReader(testSale))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	SellStock(w, r)

	// check the response
	if status := w.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusBadRequest)
	}
}

func TestSellMoreStockThanOwned(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testSale := `{"username":"eric", "ticker":"LMT", "shares":2}`

	var cred Credentials
	DB.Where("username = ?", "eric").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	// set reader and writer
	r := httptest.NewRequest("POST", "/userstock/sell", strings.NewReader(testSale))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	SellStock(w, r)

	// check the response
	if status := w.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusBadRequest)
	}

	// verify no stocks were sold
	var nofunds_ownership UserStocks
	if err := DB.Where("username = ?", "nofunds").First(&nofunds_ownership).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Username not found")
		}
	}
}

func TestGetStocksOwned(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testuser := `{"username":"john"}`

	r := httptest.NewRequest("GET", "/userstock", strings.NewReader(testuser))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetStocksOwned(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains correct parameters. see "expect"
	expect := `"username":"john","ticker":"MSFT","share":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

func TestResetAccount(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testuser := `{"username":"john"}`

	// check that the user owns stocks
	var ownedstocks []UserStocks
	DB.Find(&ownedstocks, "username = ?", "john")

	for _, st := range ownedstocks {
		fmt.Printf("Stock: %s, Quantity: %d", st.Ticker, st.Shares)
	}

	// send router and writer
	r := httptest.NewRequest("POST", "/userstock/reset", strings.NewReader(testuser))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ResetAccount(w, r)

	//check that the user no longer owns stocks according to database
	DB.Find(&ownedstocks, "username = ?", "john")
	if len(ownedstocks) != 0 {
		t.Errorf("User still owns stocks: ")
		for _, st := range ownedstocks {
			t.Errorf("Stock: %s, Quantity: %d\n", st.Ticker, st.Shares)
		}
	}

	// check that user's funds are the default ($1000)
	var cred Credentials
	DB.Where("username = ?", "john").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	if cred.Funds != 1000 {
		t.Errorf("Error reseting funds. Current funds: %f. Should be 1000", cred.Funds)
	}
}
