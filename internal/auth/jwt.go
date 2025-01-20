package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var singKey = []byte("secret")

func GenerateJWT(email string) (string, error) {

	claims := jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SignedString(singKey)

	if err != nil {
		return "", err
	}

	return strToken, nil

}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return singKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["username"].(string)
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}
