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
		"Username": "testSUfunc",
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

	// This makes sure to delete the new signup before leaving. This allows us to just run this again and again without
	// manually changing
	deleteWriter := httptest.NewRecorder()
	deleteRequest := httptest.NewRequest("POST", "/credentials/delete", strings.NewReader(testLoginCreds))
	request.Header.Set("Content-Type", "application/json")
	deleteCredentials(deleteWriter, deleteRequest)
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

func TestDelete(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "testDelete",
		"Password": "test" }`

	signUpRequest := httptest.NewRequest("POST", "/credentials/signup", strings.NewReader(testLoginCreds))
	signUpRequest.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()

	if err != nil {
		t.Fatal(err)
	}
	// We must first create a login to delete.
	signup(writer, signUpRequest)

	// This makes sure to delete the new signup before leaving. This allows us to just run this again and again without
	// manually changing
	deleteWriter := httptest.NewRecorder()
	deleteRequest := httptest.NewRequest("POST", "/credentials/delete", strings.NewReader(testLoginCreds))
	deleteRequest.Header.Set("Content-Type", "application/json")
	deleteCredentials(deleteWriter, deleteRequest)
	if status := deleteWriter.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteFail(t *testing.T) {
	InitialMigration()

	testLoginCreds := `{
		"Username": "testNotInDB",
		"Password": "test" }`

	if err != nil {
		t.Fatal(err)
	}
	// We must first create a login to delete.

	// We delete a username not in the DB
	deleteWriter := httptest.NewRecorder()
	deleteRequest := httptest.NewRequest("POST", "/credentials/signup", strings.NewReader(testLoginCreds))
	deleteRequest.Header.Set("Content-Type", "application/json")

	deleteCredentials(deleteWriter, deleteRequest)

	if status := deleteWriter.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}
