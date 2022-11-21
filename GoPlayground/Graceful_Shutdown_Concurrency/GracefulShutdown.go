package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)
	fmt.Println("Press Ctrl + C to interrupt !!!")
	go play(interruptCh)
	select {}
}

func play(interruptCh chan os.Signal) {
	for {
		select {
		case <-interruptCh:
			fmt.Println("Gracefully shutdown in progress.")
			os.Exit(1)
		}
	}
}
