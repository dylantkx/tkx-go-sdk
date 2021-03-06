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

func TestGetMarketSummaries(t *testing.T) {
	summaries, err := apiClient.MarketAPI.GetMarketSummaries()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market summary #1:", summaries[0])
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
	fmt.Println("Market sell order #1:", orders[0])
}

func TestGetMarketAvgPrice(t *testing.T) {
	market := "MYR-BTC"
	avgPrice, err := apiClient.MarketAPI.GetMarketCurrentAveragePrice(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market avg price:", *avgPrice)
}

func TestGetMarketTicker24h(t *testing.T) {
	market := "MYR-BTC"
	ticker, err := apiClient.MarketAPI.GetMarketTicker24h(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market ticker 24h:", *ticker)
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

func TestGetTradeHistory(t *testing.T) {
	market := "MYR-BTC"
	history, err := apiClient.MarketAPI.GetTradeHistory(market, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market history:", history)
}

func TestGetCandles(t *testing.T) {
	market := "MYR-BTC"
	candles, err := apiClient.MarketAPI.GetCandles(market, "1D", -1, -1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market candles:", candles)
}

func TestGetRawCandles(t *testing.T) {
	market := "MYR-BTC"
	candles, err := apiClient.MarketAPI.GetRawCandles(market, "1D", -1, -1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Market raw candles:", candles)
}
