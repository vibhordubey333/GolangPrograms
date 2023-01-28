package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	wg.Add(1)

	go printValue(ctx)

	//Cancelling signal after 4 seconds
	time.Sleep(time.Second * 4)
	cancelFunc()

	wg.Wait()
}

func printValue(ctx context.Context) {
	defer wg.Done()
	counter := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received Signal. Cancelling now.")
			return
		default:
			fmt.Println("Counter: ", counter)
			counter++
			time.Sleep(time.Second * 2)
		}
	}

}
/*
Counter:  0
Counter:  1
Counter:  2
Received Signal. Cancelling now.
*/
