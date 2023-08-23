package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var frontWebSocket *websocket.Conn

type TradeMessage struct {
	Type string `json:"type"`
	Data []struct {
		Symbol string  `json:"s"`
		Price  float64 `json:"p"`
	} `json:"data"`
}

type StockMap struct {
	Price   float64
	Counter int
}

// Handles the websocket connection to Finnhub API
func APIWebSocket(conn *websocket.Conn) {
	//formats symbols and type into json format. Sends this as a text msg in the websocket connection.
	symbols := []string{"WFC", "MSFT", "LMT", "AAPL", "KO"}
	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		conn.WriteMessage(websocket.TextMessage, msg)
	}

	// "send off" struct
	stock_map := make(map[string]StockMap)

	//receive and process messages from the websocket connection
	var msg TradeMessage
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", msg)
		// process msg
		for _, entry := range msg.Data {
			// get last price stored and compare with new price received
			mapData := stock_map[entry.Symbol]

			if mapData.Price != entry.Price {
				mapData = StockMap{
					Price:   entry.Price,
					Counter: mapData.Counter + 1,
				}
				stock_map[entry.Symbol] = mapData
				fmt.Printf("")
				// check if front web socket available, and counter = 5, to throttle data.
				if frontWebSocket != nil && mapData.Counter == 4 {
					mapData = stock_map[entry.Symbol]
					mapData.Counter = 0
					stock_map[entry.Symbol] = mapData
					// "send off" struct
					updatedMsg := TradeMessage{
						Type: "real-time-update",
						Data: []struct {
							Symbol string  `json:"s"`
							Price  float64 `json:"p"`
						}{
							{
								Symbol: entry.Symbol,
								Price:  entry.Price,
							},
						},
					}
					// update database
					st := Stock{
						Ticker: entry.Symbol,
						Price:  float32(entry.Price),
					}
					// update database
					SaveStockPrice(st)
					// send to frontend through websocket

					err := frontWebSocket.WriteJSON(updatedMsg)
					if err != nil {
						fmt.Printf("Error sending msg to front: %+v\n", err)
					}
				} else if frontWebSocket == nil {
					fmt.Println("frontWebSocket nil")
					return
				}
			}
		}
	}
}

func FrontWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("In frontwebsocket\n")
	// upgrade connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	frontWebSocket = conn
	defer func() {
		frontWebSocket.Close()
		frontWebSocket = nil
	}()

	// keep connection alive
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	}
}

func sendMockDataToFrontend() {
	ticker := time.NewTicker(5 * time.Second) // Change the interval as per your requirement
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate mock data
			price := 200.0 + rand.Float64()*100.0
			mockData := TradeMessage{
				Type: "mock-data",
				Data: []struct {
					Symbol string  `json:"s"`
					Price  float64 `json:"p"`
				}{
					{
						Symbol: "MSFT",
						Price:  price,
					},
				},
			}

			// Send the mock data to the frontend websocket
			if frontWebSocket != nil {
				err := frontWebSocket.WriteJSON(mockData)
				if err != nil {
					fmt.Printf("Error sending mock data to frontend: %+v\n", err)
				}
			} else {
				fmt.Printf("Front websocket nil")
			}
		}
	}
}
