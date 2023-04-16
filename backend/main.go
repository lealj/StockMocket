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
	// Make sure urls with {username} can't be confused with another url ie:
	// userstock/{username} and userstock/reset

	// funcs regarding credentials & user info
	//secure := rout.PathPrefix("/credentials").Subrouter()
	//secure.Use(authenticateToken)

	rout.HandleFunc("/credentials/authorize", authTokenGetClaims).Methods("GET")
	rout.HandleFunc("/credentials/signup", signup).Methods("POST")
	rout.HandleFunc("/credentials/login", login).Methods("POST")
	rout.Handle("/credentials/delete", JWTAuthProtection(http.HandlerFunc(deleteCredentials))).Methods("POST")
	rout.HandleFunc("/credentials/logout", logout).Methods("GET")
	rout.HandleFunc("/credentials/funds", GetUserFunds).Methods("POST")

	// funcs regarding what user owns
	rout.HandleFunc("/userstock/owned", GetStocksOwned).Methods("POST") //pass in username
	rout.HandleFunc("/userstock/buy", PurchaseStock).Methods("POST")    //pass in username, ticker, shares
	rout.HandleFunc("/userstock/sell", SellStock).Methods("POST")       //pass in username, ticker, shares
	rout.HandleFunc("/resetaccount", ResetAccount).Methods("POST")      //pass in username

	// portfolio funcs
	rout.HandleFunc("/portfoliohistory", GetLogs).Methods("POST")            //pass in username
	rout.HandleFunc("/portfoliovalue", GetUserPortfolioInfo).Methods("POST") //pass in username

	// funcs regarding stock and market info
	rout.HandleFunc("/stocks", GetStocks).Methods("GET")
	rout.HandleFunc("/stock", GetStock).Methods("GET")           //pass in ticker
	rout.HandleFunc("/querystocks", QueryStocks).Methods("POST") // see func

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
	UpdateStocks()
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}
