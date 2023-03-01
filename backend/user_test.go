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

func TestGetUsers(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// create a request with the mock database
	r := httptest.NewRequest("GET", "/users", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	GetUsers(w, r)

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
	expect := `"firstname":"Test","lastname":"McTest","email":"test@gmail.com","balance":0`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

func TestGetUser(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// create a request with the mock database
	r := httptest.NewRequest("GET", "/users/1", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	GetUser(w, r)

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
	expect := `"firstname":"Test","lastname":"McTest","email":"test@gmail.com","balance":0`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}

}

func TestCreateUser(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	test_user := `{"ID":2,"Firstname":"Creationtest","Lastname":"Testolgus","Email":"testagog@gmail.com","Balance":0}`
	// create a request with the mock database
	r := httptest.NewRequest("POST", "/users", strings.NewReader(test_user))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CreateUser(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//call getusers to see if user was added
	r = httptest.NewRequest("GET", "/users", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	GetUsers(w, r)

	body := w.Body.String()
	//tests if body contains correct id
	if !strings.Contains(body, `"ID":2`) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, `"ID":2`)
	}

	//tests if body contains correct parameters. see "expect"
	expect := `"firstname":"Creationtest","lastname":"Testolgus","email":"testagog@gmail.com","balance":0`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}

	//DB.Unscoped().Where("id = ?", 2).Delete(&User{})
}

func TestUpdateUser(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	test_user := `{"ID":2,"Firstname":"UpdatedFirst","Lastname":"Testolgus","Email":"testagog@gmail.com","Balance":300}`

	// create a request with the mock database
	r := httptest.NewRequest("PUT", "/users/2", strings.NewReader(test_user))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// call the function with the mock database
	UpdateUser(w, r)

	//call getusers to see if user was added
	r = httptest.NewRequest("GET", "/users", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	GetUsers(w, r)

	body := w.Body.String()
	//tests if body contains correct id
	if !strings.Contains(body, `"ID":2`) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body, `"ID":2`)
	}

	//tests if body contains correct parameters. see "expect"
	expect := `"firstname":"UpdatedFirst","lastname":"Testolgus","email":"testagog@gmail.com","balance":300`
	if !strings.Contains(body, expect) {
		t.Errorf("Unexpected body returned. got %v; want body to contain: %v", body,
			expect)
	}
}

// Doesn't work... delete works fine in actual application, just not this unit test.
func TestDeleteUser(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB

	// create a request with the mock database
	r := httptest.NewRequest("DELETE", "/users/2", nil)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	DeleteUser(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//call getusers to see if user was added
	r = httptest.NewRequest("GET", "/users", nil)
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	GetUsers(w, r)

	// check the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := w.Body.String()
	//tests if body contains correct id
	if strings.Contains(body, `"ID":2`) {
		t.Errorf("Unexpected body returned. got %v; want body to not contain: %v", body, `"ID":2`)
	}

}
