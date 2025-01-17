package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var singKey = []byte("secret")

func GenerateJWT(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24)
	claims["user"] = name
	claims["authorized"] = true

	strToken, err := token.SignedString(singKey)

	if err != nil {
		return "", err
	}

	return strToken, nil

}
