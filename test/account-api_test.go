package test

import (
	"fmt"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	info, err := apiClient.AccountAPI.GetAccountInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Account info:", info)
}

func TestGetAccountBalances(t *testing.T) {
	balances, err := apiClient.AccountAPI.GetAccountBalances()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Account balances:", balances)
}

func TestGetAccountBalanceByCCY(t *testing.T) {
	balance, err := apiClient.AccountAPI.GetAccountBalanceByCCY("BTC")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Account balance:", balance)
}

func TestGetDepositAddress(t *testing.T) {
	addr, err := apiClient.AccountAPI.GetDepositAddress("BTC")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Deposit address:", addr)
}
