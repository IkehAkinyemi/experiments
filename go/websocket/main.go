package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
    // Handle WebSocket connections
    http.Handle("/ws", websocket.Handler(handleWS))

    // Serve the HTTP server on port 8080
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func handleWS(ws *websocket.Conn) {
    // Handle incoming messages
    for {
        var msg string
        err := websocket.Message.Receive(ws, &msg)
        if err != nil {
            fmt.Println("Error receiving message:", err)
            break
        }

        fmt.Println("Received message:", msg)

        // Send a response message
        response := fmt.Sprintf("Received message: %s", msg)
        err = websocket.Message.Send(ws, response)
        if err != nil {
            fmt.Println("Error sending message:", err)
            break
        }
    }
}
