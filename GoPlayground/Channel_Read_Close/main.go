package main

import (
	"fmt"
)

func main() {

	notification := make(chan bool)
	number := make(chan int)
	go readNumber(notification, number)
	for i := 1; i <= 10; i++ {
		fmt.Print("  ", <-number)
	}
	<-notification
	fmt.Println("\nProgram exiting all go routines are completed")
}

func readNumber(notif chan bool, number chan int) {

	for i := 0; i < 10; i++ {
		number <- i
	}

	notif <- true
}
