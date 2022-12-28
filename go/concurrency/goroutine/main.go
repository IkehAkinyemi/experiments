package main

import (
	"fmt"
	"sync"
)

type acctBalance struct {
	mu     sync.Mutex
	amount int
}

func main() {

	userAcct := acctBalance{amount: 100}

	var wg sync.WaitGroup
	wg.Add(2)

	// Fund acct operation
	go func() {

		userAcct.mu.Lock()
		defer userAcct.mu.Unlock()

		userAcct.amount += 40
		wg.Done()
	}()

	// withdraw from acct operation
	go func() {
		userAcct.mu.Lock()
		defer userAcct.mu.Unlock()

		userAcct.amount -= 30
		wg.Done()
	}()

	wg.Wait()

	fmt.Print(userAcct.amount)
}
