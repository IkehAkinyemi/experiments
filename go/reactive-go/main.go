package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
    input := []int{1, 2, 3, 4, 5}

    observable := rxgo.FromSlice(input).
        Map(func(x interface{}) interface{} {
            return x.(int) * 2
        }).
        Filter(func(x interface{}) bool {
            return x.(int) > 5
        }).
        Sum()

    result, err := observable.Last()
    if err != nil {
        fmt.Printf("Error: %v", err)
    } else {
        fmt.Printf("Result: %d", result.V)
    }


		observableClient := rxgo.Just("Hello, world!").
        Map(func(x interface{}) interface{} {
            return fmt.Sprintf("%s from the client", x.(string))
        })

		observableClient.Subscribe(func(x rxgo.Item) {
        fmt.Println(x.V)
    })
}
