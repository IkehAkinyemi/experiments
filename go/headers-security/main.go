package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

const (
	apiKey    = "MY_API_KEY"
	apiSecret = "MY_API_SECRET"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the request headers
		requestApiKey := r.Header.Get("X-API-KEY")
		requestExpiry := r.Header.Get("X-EXPIRY")
		requestSignature := r.Header.Get("X-REQUEST-SIGNATURE")

		// Check that the API key is valid
		if requestApiKey != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid API key")
			return
		}

		// Check that the request has not expired
		expiryTime, err := time.Parse(time.RFC3339, requestExpiry)
		if err != nil || expiryTime.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Request has expired")
			return
		}

		// Generate the expected signature using the request data and API secret
		requestData := []byte("{" + `"param1":"value1","param2":"value2"` + "}")
		requestHash := hmac.New(sha256.New, []byte(apiSecret))
		requestHash.Write([]byte(apiKey))
		requestHash.Write([]byte(requestExpiry))
		requestHash.Write(requestData)
		expectedSignature := hex.EncodeToString(requestHash.Sum(nil))

		// Check that the request signature matches the expected signature
		if requestSignature != expectedSignature {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid request signature")
			return
		}

		fmt.Println(expectedSignature)

		// If all checks pass, return a successful response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello World!")
		fmt.Println("Sent!")
	})

	http.ListenAndServe(":8080", nil)
}
