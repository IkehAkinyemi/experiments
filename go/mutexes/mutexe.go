package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// import (
// 	"strconv"
// 	"sync"
// )

// type balance struct {
// 	amount   float64
// 	currency string
// 	mu       sync.RWMutex
// }

// var mu = &sync.Mutex{}
// var mybalance = &balance{amount: 50.00, currency: "GBP"}

// func (b *balance) Add(i float64) {
// 	b.mu.Lock()
// 	b.amount += i
// 	b.mu.Unlock()
// }

// func (b *balance) Display() string {
// 	b.mu.RLock()
// 	defer b.mu.RUnlock()
// 	return strconv.FormatFloat( b.amount, 'f', 2, 64) + " " + b.currency
// }

func main() {
	var wg sync.WaitGroup

	w := newWord()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func (filename string)  {
			defer wg.Done()

			if err := tallyword(filename, w); err != nil {
				fmt.Println(err.Error())
			}
		}(f)
	}

	wg.Wait()

	fmt.Printf("%#v\n", w)
}

type words struct {
	found map[string]int
	sync.Mutex
}

func newWord() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}

	w.found[word] = count + n
}

func tallyword(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}

	return scanner.Err()
}