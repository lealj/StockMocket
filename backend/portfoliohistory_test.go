package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLogs(t *testing.T) {
	mockDB := MockDB_Init()

	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// just purchase a stock here, test log later in func
	test_user := `{"username":"john"}`

	r := httptest.NewRequest("POST", "/portfoliohistory/john", strings.NewReader(test_user))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetLogs(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()
	// expect
	expect := `"username":"john"`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

func TestGetUserPortfolioInfo(t *testing.T) {
	mockDB := MockDB_Init()

	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// just purchase a stock here, test log later in func
	test_user := `{"username":"john"}`

	r := httptest.NewRequest("POST", "/portfoliovalue/john", strings.NewReader(test_user))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetUserPortfolioInfo(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v\n want %v\n", status, http.StatusOK)
	}

	body := w.Body.String()
	// expect
	expect := `"portfolio_value":`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}
