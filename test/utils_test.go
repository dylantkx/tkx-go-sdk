package test

import (
	"reflect"
	"testing"
	"tkx-api-client"
)

func TestTokenGenerator(t *testing.T) {
	clientID := "4A0A9B12D33F57B816EC"
	clientSecret := "secret"
	g := &tkxclient.TokenGenerator{}
	token, err := g.Generate(clientID, clientSecret)
	if err != nil {
		t.Error(err)
	}
	if reflect.ValueOf(token).Kind() != reflect.String {
		t.Error("Token is not a string")
	}
	println("Token:", token)
}
