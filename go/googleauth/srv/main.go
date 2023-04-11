package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/api/idtoken"
)

type GoogleAuthResponse struct {
  Token string `json:"token"`
}

type User struct {
  Email    string `json:"email"`
  Name     string `json:"name"`
  Picture  string `json:"picture"`
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
  var authResponse GoogleAuthResponse
  if err := json.NewDecoder(r.Body).Decode(&authResponse); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // Verify the authentication token with the Google API
  idToken, err := idtoken.Validate(context.Background(), authResponse.Token, os.Getenv("GOOGLE_CLIENT_ID"))
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // Retrieve the user's information from the token
  var user User
  user.Email = idToken.Claims["email"].(string)
  user.Name = idToken.Claims["name"].(string)
  user.Picture = idToken.Claims["picture"].(string)

  // Create a session for the user and generate a JWT
  // ...

  // Return the JWT to the frontend
  response := map[string]string{"token": token}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func main() {
  http.HandleFunc("/api/login/google", handleGoogleLogin)
  fmt.Println("Listening on :8080")
  http.ListenAndServe(":8080", nil)
}
