package tkxsdk

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenGenerator struct {
}

func (generator *TokenGenerator) Generate(clientID, clientSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "tkx",
		"id":  clientID,
	})
	tokenString, err := token.SignedString([]byte(clientSecret))
	return tokenString, err
}
