package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func httpHandler() http.Handler {
	rout := mux.NewRouter()

	/* API REQUESTS */

	// funcs regarding credentials & user info
	//secure := rout.PathPrefix("/credentials").Subrouter()
	//secure.Use(authenticateToken)

	rout.HandleFunc("/credentials/authorize", authTokenGetClaims).Methods("GET")
	rout.HandleFunc("/credentials/signup", signup).Methods("POST")
	rout.HandleFunc("/credentials/login", login).Methods("POST")
	rout.Handle("/credentials/delete", JWTAuthProtection(http.HandlerFunc(deleteCredentials))).Methods("POST")
	rout.HandleFunc("/credentials/logout", logout).Methods("GET")

	// funcs regarding what user owns
	rout.HandleFunc("/userstock/{username}", GetStocksOwned).Methods("POST")
	rout.HandleFunc("/userstock/buy/{username}", PurchaseStock).Methods("POST")
	rout.HandleFunc("/userstock/sell/{username}", SellStock).Methods("POST")

	// funcs regarding stock and market info
	rout.HandleFunc("/stocks", GetStocks).Methods("GET")

	//issues here
	rout.HandleFunc("/stocks/{ticker}", GetStock).Methods("GET")

	rout.HandleFunc("/updatestocks", UpdateStocks).Methods("GET")
	rout.HandleFunc("/querystocks", QueryStocks).Methods("POST")

	//must be last
	rout.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(rout))
}

func main() {
	host := "127.0.0.1:8080"
	InitialMigration()
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}
