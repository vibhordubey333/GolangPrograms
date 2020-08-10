package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	theOre := []string{"Iron", "Silver", "Gold", "Aluminium", "Copper"}

	oreChan := make(chan string)
	wg.Add(2)
	go func(ore []string) {

		for _, v := range ore {

			//Writing to channel.
			oreChan <- v
		}
		close(oreChan) // Very important to mention without it , program will panic
		//
		wg.Done()
	}(theOre)

	for v := range oreChan {
		fmt.Println("Value is :", v)
	}
}
