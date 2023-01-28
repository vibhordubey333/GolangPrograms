package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	chEven := make(chan int)
	chOdd := make(chan int)
	interruptCh := make(chan os.Signal, 1)

	N := 10
	//We're adding 2 as a count because anonymous goroutine where our interrupt is listening
	// is not important. It is only needed if someone terminates this program while our odd/even
	// goroutine are printing.
	wg.Add(2)

	signal.Notify(interruptCh, os.Interrupt)
	go printOddNo(chOdd, N)
	go printEvenNo(chEven, N)
	//Interrupt for handling
	go func() {
		for {
			select {
			case <-interruptCh:
				fmt.Println("Gracefully shutdown in progress.")
				os.Exit(-1)
			}
		}
	}()
	//Listening for Odd Channel goroutine.
	go func() {
		for i := range chOdd {
			fmt.Println("Odd: ", i)
		}
	}()
	//Listening for Even channel goroutine.
	go func() {
		for i := range chEven {
			fmt.Println("Even: ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
}

func printEvenNo(chEven chan int, N int) {
	defer close(chEven)
	for i := 0; i < N; i++ {
		chEven <- i
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func printOddNo(chOdd chan int, N int) {
	defer close(chOdd)

	for i := 0; i < N; i++ {
		chOdd <- i
	}

	wg.Done()
}
