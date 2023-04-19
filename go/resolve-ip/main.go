package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("8.8.8.8")
	names, err := net.LookupAddr(ip.String())
	if err != nil {
		fmt.Println("Could not resolve address:", err)
		return
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
