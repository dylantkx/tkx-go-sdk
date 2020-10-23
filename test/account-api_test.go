package test

import (
	"fmt"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	json, err := tkxClient.AccountAPI.GetAccountInfo()
	if err != nil {
		t.Error(err)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Account info:", json.Data)
}

func TestGetAccountBalances(t *testing.T) {
	json, err := tkxClient.AccountAPI.GetAccountBalances()
	if err != nil {
		t.Error(err)
	}
	if json.Data == nil {
		t.Error("Data is null")
	}
	fmt.Println("Account balances:", json.Data)
}
