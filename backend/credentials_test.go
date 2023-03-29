package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "test",
		"Password": "test" }`

	request := httptest.NewRequest("POST", "/credentials/login", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	if err != nil {
		t.Fatal(err)
	}
	login(writer, request)

	if status := writer.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLoginFalse(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "testFalse",
		"Password": "thisshouldfail" }`

	request := httptest.NewRequest("POST", "/credentials/login", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	if err != nil {
		t.Fatal(err)
	}

	login(writer, request)

	if status := writer.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestSignup(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "newsignupfafvnjabncjaib",
		"Password": "testsicaaccfngudsgsb3" }`

	request := httptest.NewRequest("POST", "/credentials/signup", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()

	if err != nil {
		t.Fatal(err)
	}

	signup(writer, request)

	if status := writer.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestSignupFail(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "test",
		"Password": "test" }`

	request := httptest.NewRequest("POST", "/credentials/signup", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()

	if err != nil {
		t.Fatal(err)
	}

	signup(writer, request)

	if status := writer.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}
