package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	privateKey, err := os.ReadFile("keys/private_key.pem")
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	fmt.Println(tokenString)

	file, err := os.Create("tokens/token.txt")
	if err != nil {
		log.Fatalf("Failed to create token file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(tokenString)
	if err != nil {
		log.Fatalf("Failed to write token to file: %v", err)
	}

	fmt.Println("Token saved to file: tokens/token.txt")
}
