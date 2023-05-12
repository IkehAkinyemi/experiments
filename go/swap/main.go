package main

import "fmt"

func swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}

func main() {
	x := 10
	s := 12

	swap(&x, &s)

	fmt.Println(x, s)
}