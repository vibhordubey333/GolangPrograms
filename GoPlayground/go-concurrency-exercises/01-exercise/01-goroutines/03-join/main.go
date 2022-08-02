package main

import (
	"fmt"
	"sync"
)

var wgObject sync.WaitGroup

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int
	wgObject.Add(1)
	go func() {
		defer wgObject.Done()
		data++
	}()
	wgObject.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
