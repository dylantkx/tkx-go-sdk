package test

import (
	"fmt"
	"testing"
)

func TestPlaceLimitBuyOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	price := 3980.0
	order, err := apiClient.OrderAPI.PlaceLimitBuyOrder(market, units, price)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order placed:", order)
}

func TestPlaceMarketBuyOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	order, err := apiClient.OrderAPI.PlaceMarketBuyOrder(market, units)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order placed:", order)
}

func TestPlaceLimitSellOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	price := 50000.0
	order, err := apiClient.OrderAPI.PlaceLimitSellOrder(market, units, price)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order placed:", order)
}

func TestPlaceMarketSellOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.01
	order, err := apiClient.OrderAPI.PlaceMarketSellOrder(market, units)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order placed:", order)
}

func TestGetMySellOrders(t *testing.T) {
	market := "MYR-BTC"
	orders, err := apiClient.OrderAPI.GetMySellOrders(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("My sell orders:", orders)
}

func TestGetMyBuyOrders(t *testing.T) {
	market := "MYR-BTC"
	orders, err := apiClient.OrderAPI.GetMyBuyOrders(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("My buy orders:", orders)
}

func TestGetMyOrders(t *testing.T) {
	market := "MYR-BTC"
	orders, err := apiClient.OrderAPI.GetMyOrders(market)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("My orders:", orders)
}

func TestCancelOrder(t *testing.T) {
	ok, err := apiClient.OrderAPI.CancelOrder(2226310)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Order is cancelled?", ok)
}
