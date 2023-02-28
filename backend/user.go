package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	//stock positions, etc.
}

func testLogin(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var userCredentials = &Credentials{}
	err := json.NewDecoder(rout.Body).Decode(&userCredentials)

	// This is a precaution so body is correct
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting users with same username that are already in the database
	result := DB.First(&userCredentials, "password = ?", &userCredentials.Password)

	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	existingCredentials := &Credentials{}
	err = result.Scan(&existingCredentials.Password).Error

	if err != nil {
		// if no user with the same username, send back code 401
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		// if error is one of any other type, send 500 status
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if the passwords don't match, we send back another 401
	if existingCredentials.Password != userCredentials.Password {
		writer.WriteHeader(http.StatusUnauthorized)
	}
	// When we get here, that means that the passwords matched

	fmt.Fprint(writer, "Request received to log in\n")
	log.Printf("The passwords matched, the status code should be 200\n")
}

func GetUsers(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []User

	fmt.Fprint(writer, "Request recieved to get users\n")
	log.Printf("Request Received to get users\n")

	//gets table info from db
	DB.Find(&users)
	//sends information
	json.NewEncoder(writer).Encode(users)
}

func GetUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	id := params["id"]

	//prints request info in http (postman)
	fmt.Fprint(writer, "Request recieved for user id: ", id, "\n")
	//prints request info in terminal
	log.Printf("Request Received for user id: %s\n", id)
	var user []User
	DB.First(&user, id)
	json.NewEncoder(writer).Encode(user)
}

func CreateUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User

	fmt.Fprint(writer, "Request recieved to create user\n")
	log.Printf("Request Received to create user\n")

	//handles data received from request (json data)
	json.NewDecoder(rout.Body).Decode(&user)
	DB.Create(&user)
	//sends info to browser
	json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	id := params["id"]

	fmt.Fprint(writer, "Request recieved to update user of id: ", id, "\n")
	log.Printf("Request Received to update user of id: %s\n", id)

	var user User
	//if this doesn't work, revert to params["id"] in second arg of DB.first
	DB.First(&user, id)
	json.NewDecoder(rout.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(writer).Encode(user)
}

func DeleteUser(writer http.ResponseWriter, rout *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(rout)
	id := params["id"]

	fmt.Fprint(writer, "Request recieved to delete user of id: ", id, "\n")
	log.Printf("Request Received to delete user of id: %s\n", id)

	var user []User
	DB.Delete(&user, params["id"])
	json.NewEncoder(writer).Encode("User deleted.")
}
