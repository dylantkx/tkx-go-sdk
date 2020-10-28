package test

import (
	"github.com/dylantkx/tkx-go-sdk"
	"github.com/joho/godotenv"
	"os"
)

var apiClient *tkxsdk.APIClient

// SetupTestEnv - Set up test environment
func SetupTestEnv() {
	godotenv.Load("../.env")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	apiBaseURL := "https://api.tokenize-my.com/public/v1"
	apiClient = tkxsdk.NewAPIClientWithCredentials(apiBaseURL, clientID, clientSecret)
}

// CleanUpAfterTests - Clean up after tests
func CleanUpAfterTests() {
	println("cleaning after testing")
}
