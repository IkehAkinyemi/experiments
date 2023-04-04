package main

import (
	"fmt"
	"log"
	"net"
)

func handleRequest(conn net.Conn) {
    defer conn.Close()

    // Get the client's IP address and port
    addr := conn.RemoteAddr().(*net.TCPAddr)
    fmt.Printf("Client IP address: %s\n", addr.IP.String())
    fmt.Printf("Client port: %d\n", addr.Port)

    // Handle the incoming request
    // ...
}

func main() {
    listener, err := net.Listen("tcp", ":8081")
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()

    fmt.Println("Listening on :8081")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go handleRequest(conn)
    }
}
