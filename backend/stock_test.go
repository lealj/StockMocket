package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mock database should already have entry: ID:1 firstname:Test lastname McTest test@gmail.com 0
func MockDB_Init() *gorm.DB {
	mock_dns := "root:cici1998@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local"
	mockDB, err := gorm.Open(mysql.Open(mock_dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open mock database: %v", err)
	}
	mockDB.AutoMigrate(&Stock{}, &UserStocks{}, &Credentials{})

	return mockDB
}

func TestGetStocks(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// create a request with the mock database
	r := httptest.NewRequest("GET", "/stocks", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	GetStocks(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := w.Body.String()
	//tests if body contains correct id
	if !strings.Contains(body, `"ID":1`) {
		t.Errorf("Unexpected body returned. got %v\n; want body to contain: %v\n", body, `"ID":1`)
	}
	//tests if body contains correct parameters. see "expect"
	expect := `"ticker":"KO","price":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v\n; want body to contain: %v\n", body, expect)
	}
}

func TestGetStock(t *testing.T) {

	mockDB := MockDB_Init()

	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	testStock := `{"ticker":"MSFT"}`

	// create a request with the mock database
	r := httptest.NewRequest("POST", "/stocks/MSFT", strings.NewReader(testStock))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	GetStock(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains correct parameters. see "expect"
	expect := `"ticker":"MSFT","price":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v\n; want body to contain: %v\n", body, expect)
	}
}

func TestUpdateStocks(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// Set prices of stocks to wrong values
	for id := 1; id <= 5; id++ {
		var tstock Stock
		DB.First(&tstock, id)
		tstock.Price = 0
		DB.Save(&tstock)
	}

	r := httptest.NewRequest("GET", "/updatestocks", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	UpdateStocks(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains the wrong price values that were set earlier ($0)
	expectedStrings := []string{`"ticker":"MSFT","price":280.51`, `"ticker":"LMT","price":474.19`,
		`"ticker":"WFC","price":37.97`, `"ticker":"AAPL","price":160.77`, `"ticker":"KO","price":61.86`}
	containedString := 0

	for _, s := range expectedStrings {
		if !strings.Contains(body, s) {
			containedString = 1
			t.Errorf("Body had incorrect prices for. %s", s)
		}
	}

	if containedString == 1 {
		t.Errorf("Here is the body received: %v", body)
	}

}

func TestQueryStocks(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	test_query := `{"ticker":"MSFT","start_month":3,"start_day":29,"start_year":2023, 
	"end_month":3,"end_day":29,"end_year":2023}`

	r := httptest.NewRequest("POST", "/querystocks", strings.NewReader(test_query))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	QueryStocks(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()

	//tests if body contains correct parameters. see "expect"
	expect := `"date":"03-29-2023","price":280.51`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}
