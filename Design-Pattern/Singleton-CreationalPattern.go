package main

import (
	"fmt"
	"sync"
)

var (
	once           sync.Once
	singleInstance *single
)

type single struct{}

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance Already Created-2")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	for {
	}
}
