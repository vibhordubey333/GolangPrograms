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
		wg.Done()
	}(theOre)

	go func() {
		for i := 0; i < len(theOre); i++ {
			//Reading from the channel.
			fmt.Println("Miner recieved ore :", <-oreChan)
		}
		wg.Done()
	}()
	wg.Wait()
}
