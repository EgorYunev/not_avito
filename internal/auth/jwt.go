package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var singKey = []byte("secret")

func GenerateJWT(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	strToken, err := token.SignedString(singKey)

	if err != nil {
		return "", err
	}

	return strToken, nil

}

func ParseJWT(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return singKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
