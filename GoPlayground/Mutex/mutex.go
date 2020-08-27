package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type GlobalRes struct {
	mux sync.Mutex
	no  int
}

func main() {
	grObj := new(GlobalRes)
	for i := 1; i <= 1000; i++ {
		//If instead of 2 you are passing 1 then it will give error negative waitgroup counter.
		//If you will pass 10 then it will give error "All goroutines are asleep - deadlock"
		wg.Add(1)

		go grObj.incValue()
		go grObj.incValue()
	}
	wg.Wait()
	fmt.Println("Final Value : ", grObj.no)
}
func (grObj *GlobalRes) incValue() {
	grObj.mux.Lock()
	defer grObj.mux.Unlock()
	grObj.no += 1
	wg.Done()
}
