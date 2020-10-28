package tkxsdk

import (
	"github.com/dgrijalva/jwt-go"
)

// TokenGenerator - TokenGenerator struct
type TokenGenerator struct {
}

// Generate - Generate a jwt token
func (generator *TokenGenerator) Generate(clientID, clientSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "tkx",
		"id":  clientID,
	})
	tokenString, err := token.SignedString([]byte(clientSecret))
	return tokenString, err
}
