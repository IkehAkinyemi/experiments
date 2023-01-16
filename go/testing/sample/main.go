package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
  flag.Parse()
  if flag.NArg() < 1 {
    fmt.Println("Usage: words-cli ")
    return
  }
  
  fsProvider := os.DirFS("/")
  
  target := flag.Arg(0)
  words, err := readWordsFromFile(fsProvider, target)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    os.Exit(1)
  }
  
  for _, word := range words {
    fmt.Println(word)
  }
}