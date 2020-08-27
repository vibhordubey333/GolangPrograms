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
			/*if i == 10 { // To create deadlock enable this.
				break
			}*/
			fmt.Println("Consumed:", <-chanInt)
		}
		//close(chanInt)
		wg.Done()
	}()
	wg.Wait()

}
func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}
func producer(ch chan int) {
	for i := 0; i < 1000; i++ {
		defer fmt.Println("Producer:")

		ch <- i
	}
	wg.Done()
}
