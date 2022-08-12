package main

import (
	"fmt"
)

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	//Set up the pipeline
	//Run the last stage of pipeline
	//Receive the values from square stage
	//Print each one until channel is closed.

	/*ch := generator(2,3)
	out := square(ch)


	for n := range out{
		fmt.Print("Output: ",n)
	}
	*/

	for n := range square(generator(2, 3)) {
		fmt.Println("Output : ", n)
	}
}
