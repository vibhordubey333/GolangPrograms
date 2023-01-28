
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	cancelCh := make(chan bool)
	
  wg.Add(1)

	go printNum(cancelCh)
	go cancelGoroutine(cancelCh)

	wg.Wait()
}

func printNum(cancelCh chan bool) {
	counter := 0

	for {

		select {
		case <-cancelCh:
			defer close(cancelCh)
			fmt.Println("Signal Received. Stopping Now.")
			return
		default:
			fmt.Println("Counter: ", counter)
			counter++
			time.Sleep(time.Second * 2)
		}
	}

	wg.Done()
}

func cancelGoroutine(cancelCh chan bool) {
	time.Sleep(time.Second * 4)
	cancelCh <- true
	wg.Done()
}
/*
Counter:  0
Counter:  1
Counter:  2
Signal Received. Stopping Now.
*/
