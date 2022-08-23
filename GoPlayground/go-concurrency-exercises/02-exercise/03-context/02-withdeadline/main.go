package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// Set deadline for goroutine to return computational result.
	//deadline := time.Now().Add(100 * time.Millisecond) // Output : work  is complete: {123}
	deadline := time.Now().Add(10 * time.Millisecond) // Output: Insufficient time given...Terminating
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	compute := func() <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			//Checking whether deadline is set or not.
			deadline, ok := ctx.Deadline()
			if ok {
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("Insufficient time given...Terminating")
					return
				}
			}
			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				fmt.Println("Work cancelled")
				return
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute()
	d, ok := <-ch
	if ok {
		fmt.Printf("work  is complete: %s\n", d)
	}
}
