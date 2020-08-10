package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type T int

func main() {

	chanInt := make(chan T)

	var i T

	for i = 0; i < 3; i++ {
		wg.Add(1)
		go printNum(chanInt, &wg)
	}

	//go func(){
	v, ok := <-chanInt
	if !ok {

		for i = 0; i < 3; i++ {
			fmt.Println(v)
		}
	}
	//  wg.Done()
	//}()

	wg.Wait()
}

func printNum(ch chan T, wg *sync.WaitGroup) {
	var i T

	for i = 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
