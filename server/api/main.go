package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	//router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	//router.HandleFunc("/users", updateUser).Methods("POST")
	//router.HandleFunc("/users", deleteUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {

	InitialMigration()
	initializeRouter()
	// router.HandleFunc("/hello-world", helloWorld)

	// Solves Cross Origin Access Issue

	//log.Fatal(srv.ListenAndServe())
}

//func helloWorld(writer http.ResponseWriter, request *http.Request) {
//
//  jsonBytes, err := utils.StructToJson(data)
//  if err != nil {
//    fmt.Print(err)
//  }
//
//  writer.Header().Set("Content-Type", "application/json")
//  writer.Write(jsonBytes)
//  return
//}
//func example(writer http.ResponseWriter, request *http.Request) {
//  writer.Header().Set("Content-Type", "application/json")
//}
