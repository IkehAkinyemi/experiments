package main

import (
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
    // Generate a new UUID
    myUUID := uuid.New()

    // Hash the UUID using SHA-256
    hash := sha256.Sum256([]byte(myUUID.String()))

    // Encode the hash value using base32 encoding
    encodedHash := base32.StdEncoding.EncodeToString(hash[:])

    // Remove any padding characters from the encoded string
    encodedHash = strings.TrimRight(encodedHash, "=")

    fmt.Printf("UUID: %s\n", myUUID)
    fmt.Printf("Encoded Hash: %s\n", encodedHash)
}
