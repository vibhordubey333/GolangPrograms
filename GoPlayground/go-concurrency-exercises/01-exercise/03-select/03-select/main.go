package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- " Hi There."
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case c1 := <-ch1:
			fmt.Println("Message: ", c1)
		default:
			fmt.Println("No Message Received")
		}

		// Do Some Processing
		fmt.Println("Processing....")
		time.Sleep(1500 * time.Millisecond)
	}
}
