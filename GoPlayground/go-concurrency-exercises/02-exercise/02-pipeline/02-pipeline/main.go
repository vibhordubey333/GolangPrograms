// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	totalTime := time.Now()
	in := generator(2, 3)

	// TODO: fan out square stage to run two instances.
	fmt.Println("In:", in)
	ch1 := square(in)
	ch2 := square(in)
	// TODO: fan in the results of square stages.
	for n := range merge(ch1, ch2) {
		fmt.Println("Output:", n)
	}
	fmt.Println("Time Taken: ", time.Since(totalTime))
}

/*
Working: It returns received only channel which is unbuffered. Inside main(), until first square() func
reads the value it will remain blocked. Then second value is read be 2nd square func().
*/

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

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	//Firing up another goroutine
	for _, c := range cs {
		go output(c)
	}
	//Waiting to complete before closing the channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
