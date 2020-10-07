package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func createToken(login string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func isValidToken(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false
	}

	return true
}
