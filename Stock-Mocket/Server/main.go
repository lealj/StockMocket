package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
  "fmt"
  "github.com/rs/cors"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/api/v1/example", exampleHandler).Methods("GET")
  r.HandleFunc("/create", create).Methods("POST")
  r.HandleFunc("/read", read).Methods("GET")
  r.HandleFunc("/update", update).Methods("PUT")
  r.HandleFunc("/delete", delete_).Methods("DELETE")


  r.HandleFunc("/hello-world", helloWorld)

  // Solves Cross Origin Access Issue
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4200"},
  })
  handler := c.Handler(r)

  srv := &http.Server{
    Handler: handler,
    Addr:    ":" + os.Getenv("PORT"),
  }

  log.Fatal(srv.ListenAndServe())
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {

type Bio struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
}

  jsonBytes, err := utils.StructToJson(data); if err != nil {
    fmt.Print(err)
  }

  writer.Header().Set("Content-Type", "application/json")
  writer.Write(jsonBytes)
  return
}

func create(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusOK)
        var human Bio
        err := json.NewDecoder(request.Body).Decode(&human)
        if err != nil {
                log.Fatalln("There was an error decoding the request body into the struct")
        }
        BioData = append(BioData, human)
        err = json.NewEncoder(writer).Encode(&human)
        if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }

}

func read(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Content-Type", "application/json")
                params := mux.Vars(request)["name"]
        for _, structs := range BioData {
                if structs.Name == name {
                        err := json.NewEncoder(writer).Encode(&structs)
                        if err != nil {
                                log.Fatalln("There was an error encoding the initialized struct")
                        }
                }
        }

}

func update(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Content-Type", "application/json")
        var human Bio
        err := json.NewDecoder(request.Body).Decode(&human)
        if err != nil {
                log.Fatalln("There was an error decoding the request body into the struct")
        }
        for index, structs := range BioData {
                if structs.Name == human.Name {
                        BioData = append(BioData[:index], BioData[index+1:]...)
                }
        }
        BioData = append(BioData, human)
        err = json.NewEncoder(writer).Encode(&human)
        if err != nil {
                log.Fatalln("There was an error encoding the initialized struct")
        }
}
