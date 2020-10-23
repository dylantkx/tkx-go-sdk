package test

import (
	"github.com/joho/godotenv"
	"os"
	"tkx-api-client"
)

var tkxClient *tkxclient.APIClient

// SetupTestEnv - Set up test environment
func SetupTestEnv() {
	godotenv.Load("../.env")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	tkxClient = tkxclient.NewAPIClientWithCredentials("MY", clientID, clientSecret)
}

// CleanUpAfterTests - Clean up after tests
func CleanUpAfterTests() {
	println("cleaning after testing")
}
