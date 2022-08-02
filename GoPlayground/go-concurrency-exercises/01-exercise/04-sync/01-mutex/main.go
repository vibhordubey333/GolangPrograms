package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	noOfCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(noOfCPU)
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	deposit := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance += amount
	}

	withdrawal := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	// Making 1000 deposit of 1Rs and withdrawal of 1Rs * 100 times. Balance remaining will be 900 RS.

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}
	wg.Wait()
	fmt.Println("Balance: ", balance)
}
