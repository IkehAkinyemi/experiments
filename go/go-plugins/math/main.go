package main

import (
	"fmt"
	"plugin"
)

// type MathFunc func(float64) float64

func main() {
	// Load the plugin
	p, err := plugin.Open("./myplugin.so")
	if err != nil {
		panic(err)
	}

	// Look up the symbol for the plugin's implementation of f(x)
	fsym, err := p.Lookup("FuncF")
	if err != nil {
		panic(err)
	}

	// Convert the symbol to a function pointer
	f, ok := fsym.(func(float64) float64)
	if !ok {
		fmt.Println("here")
		return
	}
	// v := *f
	

	// Use the plugin's implementation of f(x)
	x := 2.0
	result := f(x)
	fmt.Printf("f(%v) = %v\n", x, result)
}
