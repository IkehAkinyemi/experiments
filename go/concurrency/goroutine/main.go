package main

import (
	"fmt"
	"sync"
)

// go
func main() {
	// Declare a new WaitGroup.
	var wg sync.WaitGroup

	var employees = []string{
		"first_emply@xyz.com",
		"second_emply@xyz.com",
		"third_emply@xyz.com",
		"forth_emply@xyz.com",
		"fifth_emply@xyz.com",
	}
	var sent int16

	for _, employee := range employees {
		wg.Add(1)
		// code to run concurrently
		go func(emailAddr string) {
			defer wg.Done()

			// code to send email to employee
			sent += 1
		}(employee)
	}

	wg.Wait()
	fmt.Printf("successfully sent %d mails\n", sent)

	// continue other procedural execution
}
