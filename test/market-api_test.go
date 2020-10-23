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
	fmt.Println("Market info:", (*json.Data.Orders)[0])
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
	fmt.Println("Market info:", (*json.Data.Orders)[0])
}
