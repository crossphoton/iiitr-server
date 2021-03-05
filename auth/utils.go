package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func generateJWTToken(stu Student) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = stu.Email
	claims["name"] = stu.Name

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}
