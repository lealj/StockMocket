package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model

	Username string  `json:"username"`
	Password string  `json:"password"`
	Funds    float64 `json:"funds"`
}

func login(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// userCredentials has two field user and pass
	var userCredentials = &Credentials{}
	err := json.NewDecoder(router.Body).Decode(&userCredentials)

	// This is a precaution so body is correct
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// This works so far

	existingCredentials := &Credentials{}
	// Getting users with same username that are already in the database
	DB.Table("credentials").Select("username", "password").Where("username = ?", userCredentials.Username).Scan(&existingCredentials)

	//fmt.Fprint(writer, result)

	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
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
		return
	}
	// When we get here, that means that the passwords matched
	writer.WriteHeader(http.StatusOK)

	log.Printf("The passwords matched, the status code should be 200\n")
}

func signup(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newCredentials Credentials
	//default variable value
	newCredentials.Funds = 1000
	//handles data received from request (json data)
	err := json.NewDecoder(router.Body).Decode(&newCredentials)

	// This is a precaution so body is correct
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	existingCredentials := &Credentials{}
	// Getting users with same username that are already in the database
	DB.Table("credentials").Select("username", "password").Where("username = ?", newCredentials.Username).Scan(&existingCredentials)
	//check if new signup username already exists
	if existingCredentials.Username == newCredentials.Username {
		log.Println("Username already exists")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	//create new user as username doesn't exist already
	DB.Create(&newCredentials)

	fmt.Fprint(writer, "Successfully saved username and password")
	json.NewEncoder(writer).Encode(newCredentials)
}
