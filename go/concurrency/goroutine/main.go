package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// go
func main() {
	var errNotFound = errors.New("db record doesn't exist")
	var statusCode = 404

	// goroutine defintion syntax using closure.
	go func() {
		// code to run concurrently
		fmt.Fprintf(os.Stdout, "db lookup error: %v with status code %d", errNotFound, statusCode)
	}()

	time.Sleep(1 * time.Second)
}
