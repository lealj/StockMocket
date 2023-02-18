package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

//Code

var dataBase *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	dataBase, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	dataBase.AutoMigrate(&User{})
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User

	if err != nil {
		fmt.Println(err.Error())
	}
	// taking data from request body
	// of the data we have is converted into the data of the type user
	json.NewDecoder(request.Body).Decode(&user)
	dataBase.Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user User
	dataBase.First(&user, params["id"])
	json.NewEncoder(writer).Encode(user)

}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []User
	dataBase.Find(&users)
	json.NewEncoder(writer).Encode(users)

}
