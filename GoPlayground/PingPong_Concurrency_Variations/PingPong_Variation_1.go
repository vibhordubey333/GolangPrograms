package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan int, 1)
	pongChan := make(chan int, 1)

	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)

	//Ping channel doing service.
	pingChan <- 1

	//To indefinite run our program.
	select {}
	//for {} //Alternative way to indefinite run program other than select{}
}

func ping(pingChan <-chan int, pongChan chan<- int) {
	for {
		ball := <-pingChan
		fmt.Println("Ping: ", ball)
		time.Sleep(1 * time.Second)

		pongChan <- ball + 1
	}
}

func pong(pongChan <-chan int, pingChan chan<- int) {
	for {
		ball := <-pongChan
		fmt.Println("Pong: ", ball)
		time.Sleep(1 * time.Second)
		pingChan <- ball
	}
}
