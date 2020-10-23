package test

import (
	"fmt"
	"testing"
)

func TestGetMarketBuyOrders(t *testing.T) {
	market := "MYR-BTC"
	json, err := apiClient.MarketAPI.GetMarketBuyOrders(market)
	if err != nil {
		t.Error(err)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Market buy order #1:", (*json.Data.Orders)[0])
}

func TestGetMarketSellOrders(t *testing.T) {
	market := "MYR-BTC"
	json, err := apiClient.MarketAPI.GetMarketSellOrders(market)
	if err != nil {
		t.Error(err)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Market sell order #1", (*json.Data.Orders)[0])
}

func TestGetMarketOrderBook(t *testing.T) {
	market := "MYR-BTC"
	json, err := apiClient.MarketAPI.GetMarketOrderBook(market, "both", 5)
	if err != nil {
		t.Error(err)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Order book buy order #1:", (*json.Data.Buy)[0])
	fmt.Println("Order book sell order #1:", (*json.Data.Sell)[0])
}
