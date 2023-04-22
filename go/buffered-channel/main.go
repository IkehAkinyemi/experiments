package main

import (
	"fmt"
	"os"
)

func writeToFile(filename string, ch chan string) {
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("Error creating file: %s", err)
        return
    }
    defer file.Close()

    for message := range ch {
        _, err = file.WriteString(message + "\n")
        if err != nil {
            fmt.Printf("Error writing to file: %s", err)
            return
        }
    }
}

func main() {
    ch := make(chan string, 2)

    go writeToFile("output.txt", ch)

    ch <- "Hello, world!"
    ch <- "This is a test message."
    ch <- ""

    fmt.Println("File written successfully!")
}
