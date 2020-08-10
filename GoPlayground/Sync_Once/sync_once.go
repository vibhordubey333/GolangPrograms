package main

import (
	"fmt"
	"sync"
)

var syncOnceObj sync.Once

func main() {

	DoOnce()
	DoOnce()
}

func DoOnce() {
	syncOnceObj.Do(func() {
		fmt.Println("I am one of a kind.")
	})
	fmt.Println("I am omnipresent.")
}
