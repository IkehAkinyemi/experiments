package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

const (
	randomValueLength = 43 // Length of the desired random value
)

func generateRandomValue() (string, error) {
	randomBytes := make([]byte, (randomValueLength+3)/4*3)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomValue := base64.RawURLEncoding.EncodeToString(randomBytes)
	randomValue = randomValue[:randomValueLength]
	return randomValue, nil
}

func main() {
	randomValue, err := generateRandomValue()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Random Value:", randomValue)
}
