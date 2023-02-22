package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	wg.Add(2)
	go func() {
		fmt.Println(http.ListenAndServe("localhost:7777", nil))
	}()
	go appendOperation()
	wg.Wait()

}

func appendOperation() {
	wg.Done()
	fmt.Println("Start: ", time.Now())

	junkSlice := []string{}

	for i := 0; i < 99999; i++ {
		junkSlice = append(junkSlice, "01")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Stop:", time.Now())

}
