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
	// expect, dynamic variables shouldn't fail this
	expected := []string{`"username":"john","ticker":"MSFT","shares":`,
		`"ordertype":"buy","price":`}

	for _, str := range expected {
		if !strings.Contains(body, str) {
			t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, str)
		}
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
	// expect - numbers may need adjustment
	expect := []string{`"portfolio_value":1000`, `"pv_change":0`}
	for _, str := range expect {
		if !strings.Contains(body, str) {
			t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
				str)
		}
	}
}
