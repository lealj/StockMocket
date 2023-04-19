package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// JWT Authentication Testing
func TestExtractClaims(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "test",
		"Password": "test" }`

	if err != nil {
		t.Fatal(err)
	}

	// We must first log in to do anything
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/credentials/login", strings.NewReader(testLoginCreds))

	request.Header.Set("Content-Type", "application/json")
	cookie, err := request.Cookie("loggedIn")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, err := extractClaims(cookie.Value)

	if claims.Username != "test" {
		t.Errorf("handler returned wrong status code: got %v want %v", claims.Username, "test")
	}
}

func TestGetClaimsHandler(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "test",
		"Password": "test" }`

	if err != nil {
		t.Fatal(err)
	}

	// We log in to make sure that user information can be accessed.
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/credentials/login", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")

	login(writer, request)
	getClaimsHandler(writer, request)

	if status := writer.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestJWTPathProtection(t *testing.T) {
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

	dRequest := httptest.NewRequest("POST", "/credentials/delete", strings.NewReader(testLoginCreds))
	deleteCredentials(writer, dRequest)

	if status := writer.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestJWTPathProtectionFail(t *testing.T) {
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
	dRequest := httptest.NewRequest("POST", "/credentials/delete", strings.NewReader(testLoginCreds))
	deleteCredentials(writer, dRequest)
	deleteCredentials(writer, dRequest)

	if status := writer.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
