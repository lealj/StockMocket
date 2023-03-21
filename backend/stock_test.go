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
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, `"ID":1`)
	}
	//tests if body contains correct parameters. see "expect"
	expect := `"companyname":"Coca Cola Co","ticker":"KO","price":"59.23","date":"02/08/2023"`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, expect)
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

	// create a request with the mock database
	r := httptest.NewRequest("GET", "/stocks/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	GetStock(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := w.Body.String()
	//tests if body contains correct id
	if !strings.Contains(body, `"ID":1`) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, `"ID":1`)
	}
	//tests if body contains correct parameters. see "expect"
	expect := `"companyname":"Coca Cola Co","ticker":"KO","price":"59.23","date":"02/08/2023"`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, expect)
	}
}
