package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	pingCh := make(chan struct{}, 1)
	pongCh := make(chan struct{}, 1)
	interruptCh := make(chan os.Signal, 1)
	//Player 1 serve
	pingCh <- struct{}{}
	signal.Notify(interruptCh, os.Interrupt)
	go playPingPong(pingCh, pongCh, interruptCh)
	select {}
}

func playPingPong(ping, pong chan struct{}, interruptCh chan os.Signal) {
	for {
		select {
		case <-ping:
			fmt.Println("Player 1: Ping")
			time.Sleep(1 * time.Second)
			pong <- struct{}{}

		case <-pong:
			fmt.Println("Player 2: Pong")
			time.Sleep(1 * time.Second)
			ping <- struct{}{}
		case <-interruptCh:
			fmt.Println("Gracefully shutdown in progress.")
			os.Exit(1)
		}
	}
}
