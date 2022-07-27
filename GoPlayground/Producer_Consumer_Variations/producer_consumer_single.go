package main

import (
	"fmt"
	"sync"
)

var messages = []string{"A", "B", "C", "D"}

func main() {
	var wg sync.WaitGroup
	link := make(chan string)
	wg.Add(1)
	go producer(link)
	go consumer(link, &wg)
	wg.Wait()
}

func producer(link chan<- string) {
	for _, m := range messages {
		link <- m
	}
	close(link) //Channel is always closed from sender.
}

func consumer(link <-chan string, wg *sync.WaitGroup) {
	for m := range link {
		fmt.Println(m)
	}
	wg.Done()
}
