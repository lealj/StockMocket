package main

import (
	"testing"
)

func TestPurchaseStock(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB
}

func TestSellStock(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB
}

func TestGetStocksOwned(t *testing.T) {
	mockDB := MockDB_Init()
	defer func() {
		dbInstance, _ := mockDB.DB()
		_ = dbInstance.Close()
	}()

	//overwrite db variable in user.go
	DB = mockDB
}
