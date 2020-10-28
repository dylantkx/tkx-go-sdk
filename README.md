# Tokenize Public API Client

## 1. Overview
A client for Tokenize Public API implemented in Golang

**Note:** This package is under development. Use at your own risk.

## 2. Running tests
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

## 3. Contact developers/maintainers: 
- dylan@tokenize.exchange