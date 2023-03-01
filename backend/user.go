package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	//stock positions, etc.
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
