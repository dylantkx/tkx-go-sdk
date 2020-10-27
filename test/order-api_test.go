package test

import (
	"fmt"
	"testing"
)

func TestPlaceLimitBuyOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	price := 3980.0
	json, err := apiClient.OrderAPI.PlaceLimitBuyOrder(market, units, price)
	if err != nil {
		t.Error(err)
	}
	if json.Status != "success" {
		t.Error(json.Message)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Order placed:", json.Data)
}

func TestPlaceMarketBuyOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	json, err := apiClient.OrderAPI.PlaceMarketBuyOrder(market, units)
	if err != nil {
		t.Error(err)
	}
	if json.Status != "success" {
		t.Error(json.Message)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Order placed:", json.Data)
}

func TestPlaceLimitSellOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.1
	price := 50000.0
	json, err := apiClient.OrderAPI.PlaceLimitSellOrder(market, units, price)
	if err != nil {
		t.Error(err)
	}
	if json.Status != "success" {
		t.Error(json.Message)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Order placed:", json.Data)
}

func TestPlaceMarketSellOrder(t *testing.T) {
	market := "MYR-BTC"
	units := 0.01
	json, err := apiClient.OrderAPI.PlaceMarketSellOrder(market, units)
	if err != nil {
		t.Error(err)
	}
	if json.Status != "success" {
		t.Error(json.Message)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Order placed:", json.Data)
}
