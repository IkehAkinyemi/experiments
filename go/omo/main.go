package main

import "fmt"

type f struct {
	gin string
	bread string
}

func main() {
	// m := map[string]string{"gin": "holy", "bread": "water"}
	m1 := f{gin: "holy", bread: "water"}

	fio(m1, "gin", "bread")
	fio(m1, "bread", "butter")

	fmt.Printf("%+v\n", m1)
}

func tio(m map[string]string, key string, val string) {
	m[key] = val
}

func fio(m f, key string, val string) {
	m.gin = key
	m.bread = val
}
