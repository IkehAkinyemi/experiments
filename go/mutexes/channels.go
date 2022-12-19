package main

import (
	"fmt"
	"os"
	"time"
)

func cmd() {
	echo := make(chan []byte)
	timeOut := time.After(5 * time.Second)
	go echoStdin(echo)
	for {
		select {
		case out := <-echo:
			os.Stdout.Write(out)
		case <-timeOut:
			os.Exit(0)
		}
	}
}

func echoStdin(echo chan<- []byte) {
	for {
		input := make([]byte, 1024)
		n, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Errorf("%w", err)
			continue
		}
		if n > 0 {
			echo <- input
		}

	}
}
