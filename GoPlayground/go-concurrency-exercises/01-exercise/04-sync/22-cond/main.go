package main

import (
	"fmt"
	"sync"
)

var sharedRes = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()
		//Suspend goroutine until sharedRsc is populated
		for len(sharedRes) < 1 {
			c.Wait()
		}
		fmt.Println(sharedRes["rsc1"])
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		//Suspend goroutine unitl sharedRsc is populated
		c.L.Lock()
		for len(sharedRes) < 2 {
			c.Wait()
		}
		fmt.Println(sharedRes["rsc2"])
		c.L.Unlock()
	}()
	c.L.Lock()

	//Write changes to sharedRsc
	sharedRes["rsc1"] = "foo"
	sharedRes["rsc2"] = "bar"
	//Sending broadcast signal
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
