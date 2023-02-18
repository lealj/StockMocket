package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Book struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
}

func initializeRouter() {
  router := mux.NewRouter()

  router.HandleFunc("/users", getUsers).Methods("GET")
  router.HandleFunc("/users/{id}", getUser).Methods("GET")
  router.HandleFunc("/users", createUser).Methods("POST")
  // router.HandleFunc("/users", updateUser).Methods("POST")
  //router.HandleFunc("/users", deleteUser).Methods("POST")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
  InitialMigration()
  initializeRouter()
  // router.HandleFunc("/hello-world", helloWorld)

  // Solves Cross Origin Access Issue
  //c := cors.New(cors.Options{
  //  AllowedOrigins: []string{"http://localhost:4200"},
  //})
  //handler := c.Handler(r)
  //
  //srv := &http.Server{
  //  Handler: handler,
  //  Addr:    ":" + os.Getenv("PORT"),
  //}
  //
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

func create(writer http.ResponseWriter, request *http.Request) {

}

func read(writer http.ResponseWriter, request *http.Request) {

}

func update(writer http.ResponseWriter, request *http.Request) {

}

func delete_(writer http.ResponseWriter, request *http.Request) {

}

func Run() {
  router := mux.NewRouter()
  router.HandleFunc("/create", create).Methods("POST")
  router.HandleFunc("/read", read).Methods("GET")
  router.HandleFunc("/update", update).Methods("PUT")
  router.HandleFunc("/delete", delete_).Methods("DELETE")

  err := http.ListenAndServe(":8080", router)
  if err != nil {
    //log.Fatalln("There's an error with the server,", err)
  }

}
