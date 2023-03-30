package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

	r := httptest.NewRequest("POST", "/userstock/buy/john", strings.NewReader(testPurchase))
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

func TestSellStock(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testSale := `{"username":"john", "ticker":"MSFT", "shares":1}`

	var cred Credentials
	DB.Where("username = ?", "john").First(&cred)
	fmt.Printf("Funds before purchase: %f", cred.Funds)

	r := httptest.NewRequest("POST", "/userstock/sell/john", strings.NewReader(testSale))
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

func TestGetStocksOwned(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testuser := `{"username":"john"}`

	r := httptest.NewRequest("GET", "/userstock/john", strings.NewReader(testuser))
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
