package main

import (
	"fmt"
)

func main() {
	producer := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}

		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		//Read values from channels
		for v := range ch {
			fmt.Printf("Received : %d\n", v)
		}
		fmt.Println("Done receiving ")
	}
	//Calling producer and consumer.
	ch := producer()
	consumer(ch)
}
