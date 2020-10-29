package test

import (
	"fmt"
	"testing"
)

func TestGetMarkets(t *testing.T) {
	markets, err := apiClient.MarketAPI.GetMarkets()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market #1:", markets[0])
}

func TestGetLastMarketPrice(t *testing.T) {
	market := "MYR-BTC"
	lastPrice, err := apiClient.MarketAPI.GetLastMarketPrice(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Last price:", *lastPrice)
}

func TestGetMarketTicker(t *testing.T) {
	market := "MYR-BTC"
	ticker, err := apiClient.MarketAPI.GetMarketTicker(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market ticker:", *ticker)
}

func TestGetMarketBuyOrders(t *testing.T) {
	market := "MYR-BTC"
	orders, err := apiClient.MarketAPI.GetMarketBuyOrders(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market buy order #1:", orders[0])
}

func TestGetMarketSellOrders(t *testing.T) {
	market := "MYR-BTC"
	orders, err := apiClient.MarketAPI.GetMarketSellOrders(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market sell order #1", orders[0])
}

func TestGetMarketOrderBook(t *testing.T) {
	market := "MYR-BTC"
	orderbook, err := apiClient.MarketAPI.GetMarketOrderBook(market, "both", 5)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order book buy order #1:", orderbook.Buy[0])
	fmt.Println("Order book sell order #1:", orderbook.Sell[0])
}
