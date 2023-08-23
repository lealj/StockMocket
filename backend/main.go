package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func httpHandler() http.Handler {
	rout := mux.NewRouter()

	/* API REQUESTS */
	// Make sure urls with {username} can't be confused with another url ie:
	// userstock/{username} and userstock/reset

	// funcs regarding credentials & user info
	//secure := rout.PathPrefix("/credentials").Subrouter()
	//secure.Use(authenticateToken)

	rout.HandleFunc("/credentials/authorize", getClaimsHandler).Methods("GET")
	rout.HandleFunc("/credentials/signup", signup).Methods("POST")
	rout.HandleFunc("/credentials/login", login).Methods("POST")
	// A1: The delete function is protected and cannot be accessed unless the user is logged in.
	rout.Handle("/credentials/delete", JWTPathProtection(http.HandlerFunc(deleteCredentials))).Methods("POST")
	rout.HandleFunc("/credentials/logout", logout).Methods("GET")
	rout.Handle("/credentials/funds", JWTPathProtection(http.HandlerFunc(GetUserFunds))).Methods("POST")

	// user functions
	rout.HandleFunc("/userstock/owned", GetStocksOwned).Methods("POST") //pass in username
	rout.HandleFunc("/userstock/buy", PurchaseStock).Methods("POST")    //pass in username, ticker, shares
	rout.HandleFunc("/userstock/sell", SellStock).Methods("POST")       //pass in username, ticker, shares
	rout.HandleFunc("/resetaccount", ResetAccount).Methods("POST")      //pass in username

	// portfolio
	rout.HandleFunc("/portfoliohistory", GetLogs).Methods("POST")            //pass in username
	rout.HandleFunc("/portfoliovalue", GetUserPortfolioInfo).Methods("POST") //pass in username

	// stock and market info
	rout.HandleFunc("/stocksdata", GetStocks).Methods("GET")
	rout.HandleFunc("/stock", GetStock).Methods("GET")           //pass in ticker
	rout.HandleFunc("/querystocks", QueryStocks).Methods("POST") // see func

	// websocket
	rout.HandleFunc("/ws", FrontWebSocket)

	//must be last
	rout.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "X-Requested-With", "If-Modified-Since",
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

	// establish connection to 3rd party websocket
	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token=ci87gu9r01qnrgm32fggci87gu9r01qnrgm32fh0", nil)
	if err != nil {
		fmt.Printf("Error with 3rd party websocket\n")
		panic(err)
	}
	defer w.Close()
	go APIWebSocket(w)
	//go sendMockDataToFrontend()

	// establish connection to frontend websocket
	go func() {
		frontWebSocket, _, err = websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
		if err != nil {
			fmt.Printf("Error in main front web sock %+v\n", err)
		}
		defer frontWebSocket.Close()
		for {
			fmt.Printf("In main go func()\n")
			_, _, err := frontWebSocket.ReadMessage()
			if err != nil {
				fmt.Printf("Error reading msg: %+v\n", err)
				return
			}
		}
	}()

	// set up and start http server
	server := &http.Server{
		Addr:           host,
		Handler:        httpHandler(),
		MaxHeaderBytes: 64 * 1024,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}

/*
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	// Start a goroutine to execute UpdateStocks() periodically

		go func() {
			for range ticker.C {
				UpdateStocks()
			}
		}()
*/
