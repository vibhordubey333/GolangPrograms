package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	chanInt := make(chan int)
	wg.Add(2)
	go producer(chanInt)
	go func() {
		for i := 0; i < 1000; i++ {
			if i == 10 {
				break
			}
			fmt.Println("Consumed:", <-chanInt)

		}
		close(chanInt)
		wg.Done()
	}()
	wg.Wait()

}
func producer(ch chan int) {
	for i := 0; i < 1000; i++ {
		fmt.Println("Producer:")
		ch <- i
	}
	wg.Done()
}
