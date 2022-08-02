package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var initialize sync.Once
	load := func() {
		fmt.Println("Initializing Program....")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			initialize.Do(load)
		}()
	}
	wg.Wait()
}
