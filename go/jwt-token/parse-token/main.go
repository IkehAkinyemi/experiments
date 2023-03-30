package main

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	publicKey, err := os.ReadFile("keys/public_key.pem")
	if err != nil {
		panic(err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		panic(err)
	}

	tokenBytes, err := os.ReadFile("tokens/token.txt")
	if err != nil {
		panic(err)
	}
	tokenString := string(tokenBytes)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		panic("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	fmt.Println(claims["sub"], claims["name"], claims["iat"], claims["exp"])
}
