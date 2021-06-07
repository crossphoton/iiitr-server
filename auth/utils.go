package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken generated a JWT for student authentication signed in from OAuth.
func GenerateJWTToken(stu Student) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = stu.Email
	claims["name"] = stu.Name

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}

// VerifyJWT verify user using jwt token cookie.
func VerifyJWT(token string, key string) bool {

	//Decode token
	tokenDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding")
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

// GetClaims returns claims from a jwt token
func GetClaims(token string) (jwt.MapClaims, error) {

	//Decode token
	tokenDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding")
		}
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		return nil, errors.New("couldn't parse token")
	}

	if tokenDecoded.Valid {
		return tokenDecoded.Claims.(jwt.MapClaims), nil
	}

	return nil, errors.New("couldn't get claims")
}
