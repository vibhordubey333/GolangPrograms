package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Channel 1"
	}()

	select {
	case c1 := <-ch1:
		fmt.Println("Who ?", c1)
	case <-time.After(3 * time.Second):
		fmt.Println("Time out")
	}
}
