package studentdata

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func verifyJWT(token string, key string) bool {
	tokenDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error decoding")
		}
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		return false
	}

	if tokenDecoded.Valid && tokenDecoded.Claims.(jwt.MapClaims)["email"] == key {
		return true
	}

	return false
}
