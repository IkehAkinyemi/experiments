package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// EncodeNoPrefix encodes b as a hex string without 0x prefix.
func EncodeNoPrefix(b []byte) string {
	return hex.EncodeToString(b)
}
// Sha1File performs a sha1 hash on a file
func Sha1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// nolint:gosec
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("failed to copy file content to sha1 handler: %w", err)
	}

	return EncodeNoPrefix(h.Sum(nil)), nil
}

func main() {
	output, err := Sha1File("./hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(output)
}