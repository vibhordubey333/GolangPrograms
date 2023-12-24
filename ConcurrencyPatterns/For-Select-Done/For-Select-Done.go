package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go PeriodicTask(ctx)

	// Create a channel to receive signals from the operating system.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	// The code blocks until a signal is received (e.g. Ctrl+C).
	<-sigCh
}

// someTask function that we call periodically.
func someTask() {
	fmt.Println("Random No Generated : ", rand.Int()*rand.Int())
}

// PeriodicTask runs someTask every 1 second.
// If canceled goroutine should be stopped.
func PeriodicTask(ctx context.Context) {
	// Create a new ticker with a period of 1 second.
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			someTask()
		case <-ctx.Done():
			fmt.Println("stopping PeriodicTask")
			ticker.Stop()
			return
		}
	}
}
