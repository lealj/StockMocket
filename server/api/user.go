package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Code

type User struct {
	dataBase *gorm.DB

	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	var dataBase *gorm.DB
	dataBase, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("migration error")
	}
	dataBase.AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	dataBase, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	writer.Header().Set("Content-Type", "application/json")
	var user User

	// taking data from request body
	// of the data we have is converted into the data of the type user
	json.NewDecoder(request.Body).Decode(&user)
	dataBase.Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func getUsers(writer http.ResponseWriter, request *http.Request) {

}
