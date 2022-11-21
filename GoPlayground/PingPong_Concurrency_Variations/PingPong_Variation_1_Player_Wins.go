package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan string, 1)
	pongChan := make(chan string, 1)

	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)

	//Player 1 doing service.
	pingChan <- "Player 1"
	select {}
}

func ping(pingChan <-chan string, pongChan chan<- string) {
	for {
		ball := <-pingChan
		fmt.Println("Ping: ", ball)
		time.Sleep(1 * time.Second)
		pongChan <- "Player 2"
	}
}

func pong(pongChan <-chan string, pingChan chan<- string) {
	for {
		ball := <-pongChan
		fmt.Println("Pong: ", ball)
		time.Sleep(1 * time.Second)
		pingChan <- "Player 1"
	}
}
