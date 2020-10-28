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

	apiClient = tkxsdk.NewAPIClientWithCredentials("MY", clientID, clientSecret)
}

// CleanUpAfterTests - Clean up after tests
func CleanUpAfterTests() {
	println("cleaning after testing")
}
