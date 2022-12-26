package main

import (
	"fmt"
	"sync"
)

var (
	producerCount = 4
	wgObject      = sync.WaitGroup{}
	//Using 2D array to mock that each producer has set of messages to send to the consumer.
	messages = [][]string{
		{"A", "B", "C"},
		{"A1", "B1", "C1"},
		{"A2", "B2", "C2"},
		{"A3", "B3", "C3"},
	}
	wgConsumerObject = sync.WaitGroup{}
)

func main() {
	link := make(chan string)
	wgConsumerObject.Add(1)
	wgObject.Add(producerCount)
	for i := 0; i < producerCount; i++ {
		go producer(i, &wgObject, link)
	}

	go consumer(link, &wgConsumerObject)
	wgObject.Wait()
	close(link) // Closing in main groutine as multiple threads are there.
	wgConsumerObject.Wait()
}

func producer(worker int, wgObject *sync.WaitGroup, link chan<- string) {
	defer wgObject.Done()
	for _, m := range messages[worker] {
		fmt.Printf("Worker %v : Produced Message :%v\n", worker, m)
		link <- m
	}
	//Not closing link channel as multiple goroutines are firedup.
}

func consumer(link <-chan string, wgConsumerObject *sync.WaitGroup) { //done chan<- bool) {
	for m := range link {
		fmt.Println("Consumed: ", m)
	}
	wgConsumerObject.Done()
}
