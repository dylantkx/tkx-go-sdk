# Tokenize Public API Client

## 1. Overview
A client for Tokenize Public API implemented in Golang

**Note:** This package is under development. Current stable version is v0.2.0

## 2. How to use

Example usage:
```go
  func main() {
    clientID := os.Getenv("CLIENT_ID")
    clientSecret := os.Getenv("CLIENT_SECRET")
    apiBaseURL := "https://api.tokenize-my.com/public/v1"
    
    // Initialize client
    client := tkxsdk.NewAPIClientWithCredentials(apiBaseURL, clientID, clientSecret)

    // Get account info
    info, err := client.AccountAPI.GetAccountInfo()
    if err != nil {
      t.Error(err)
    }
    fmt.Println("Account info:", info)
  }
```

## 3. Running tests
To run all test cases:
```
go test ./test/ -v
```

To run a single test case:
```
go test ./test/ -run TestFunctionX -v
```

To run test with coverage:
```
go test ./test/ -v --coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html
```

## 4. Contact developers/maintainers: 
- dylan@tokenize.exchange