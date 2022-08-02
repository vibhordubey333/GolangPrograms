package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("Custom- Goroutine-1")
	// goroutine with anonymous function
	go func() {
		fun("Custom-Goroutine-2")
	}()
	// goroutine with function value call
	funVar := fun
	go funVar("Custom-Goroutine-3")

	// wait for goroutines to end
	fmt.Println("Waiting for goroutine to finish...")
	time.Sleep(1 * time.Second)
	fmt.Println("done..")
}
