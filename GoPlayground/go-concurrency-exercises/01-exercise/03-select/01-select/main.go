package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "Channel 1"
		time.Sleep(5 * time.Second)

	}()

	go func() {
		ch2 <- "Channel 2"
		time.Sleep(3 * time.Second)
	}()
	for i := 0; i < 2; i++ { //Using for loop as we want to print both the messages.
		select { //If for loop is not used whicever channel returns first then select will exit
		case c1 := <-ch1:
			fmt.Println("Who ? : ", c1)
		case c2 := <-ch2:
			fmt.Println("Who ?: ", c2)
		}
	}
}
