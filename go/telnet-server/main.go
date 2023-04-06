package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Fprintf(conn, "You wrote: %s\n", text)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading from connection: %s\n", err.Error())
    }
}

func main() {
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err.Error())
        return
    }

    defer listener.Close()

    fmt.Println("Server started on localhost:8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Error accepting connection: %s\n", err.Error())
            continue
        }
        fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
        go handleConnection(conn)
    }
}
