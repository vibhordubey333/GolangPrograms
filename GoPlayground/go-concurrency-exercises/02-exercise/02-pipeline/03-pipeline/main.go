package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Purpose of this program is to showcase where we're processing only one input and cancelling all other goroutines.
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}

	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	// To call wg.Done on all return paths.
	defer wg.Done()
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}

			out <- n
		}

	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	//Using empty struct as we don't want to send any data just for sending cancel signal.
	//We need to pass done channel to every stage.
	done := make(chan struct{})
	defer close(done)

	in := generator(done, 2, 3)

	c1 := square(done, in)
	c2 := square(done, in)

	//It is Go Idiom to pass done channel as first arguement.
	out := merge(done, c1, c2)

	// TODO: cancel goroutines after receiving one value.
	fmt.Println("Result:", <-out)

	//Closing done channel

	//It will take some time for goroutines to exit so after 10ms we will check.
	time.Sleep(10 * time.Millisecond)
	// To check if goroutines have exited
	fmt.Println("No. Of GoRoutines active: ", runtime.NumGoroutine())

}
