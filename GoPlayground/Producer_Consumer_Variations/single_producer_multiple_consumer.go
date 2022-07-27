package main

import (
	"fmt"
	"sync"
)

var (
	wgObject      sync.WaitGroup
	messages      = []string{"A", "B", "C", "A1", "B1", "C1", "A2", "B2", "C2", "A3", "B3", "C3", "A4", "B4", "C4"}
	consumerCount = 3
)

func main() {
	link := make(chan string)

	go producer(link)
	for i := 0; i < consumerCount; i++ {
		wgObject.Add(1)
		go consumer(i, link, &wgObject)
	}
	wgObject.Wait()

}

func producer(link chan<- string) {
	for _, m := range messages {
		link <- m
	}
	close(link) // Channel is closed from senders side.
}

func consumer(worker int, link <-chan string, wg *sync.WaitGroup) {
	for m := range link {
		fmt.Printf("Message %v is cosumed by worker %v \n", m, worker)
	}
	wg.Done()
}
