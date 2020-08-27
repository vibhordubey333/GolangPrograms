package main

import (
	"fmt"
	"sync"
)

//var x int = 0
type Global struct {
	mux sync.Mutex
	x   int
}

func main() {

	obj := Global{}

	//ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		//go obj.increment(ch)
		go getNumber()
	}
	fmt.Println("Final value :", obj.x)
}

func getNumber() int {
	var i int
	// Create a channel to push an empty struct to once we're done
	done := make(chan struct{})
	go func() {
		i = 5
		// Push an empty struct once we're done
		done <- struct{}{}
	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done
	return i
}

/*
func (obj *Global) increment(ch chan bool) {
	obj.mux.Lock()
	defer obj.mux.Unlock()
	ch <- true
	obj.x += 1
	<-ch
}
*/
