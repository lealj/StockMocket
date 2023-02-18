package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Bio struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var human Bio
var BioData = make([]Bio, 0)

func main() {
	router := mux.NewRouter()

	//router.HandleFunc("/api/v1/example", exampleHandler).Methods("GET")
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/read", read).Methods("GET")
	router.HandleFunc("/update", update).Methods("PUT")
	router.HandleFunc("/delete", delete_).Methods("DELETE")

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

func delete_(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)["name"]
	indexChoice := 0
	for index, structs := range BioData {
		if structs.Name == name {
			indexChoice = index
		}
	}
	BioData = append(BioData[:indexChoice], BioData[indexChoice+1:]...)
}

func Run() {
  router := mux.NewRouter()
  router.HandleFunc("/create", create).Methods("POST")
  router.HandleFunc("/read", read).Methods("GET")
  router.HandleFunc("/update", update).Methods("PUT")
  router.HandleFunc("/delete", delete_).Methods("DELETE")

  err := http.ListenAndServe(":8080", router)
  if err != nil {
    log.Fatalln("There's an error with the server," err)
  }

}
