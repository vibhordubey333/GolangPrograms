package main

import (
	"fmt"
	"sync"
)

/*
Fan out is used when multiple functions read from the same channel. The reading will stop only when the channel is closed.
This characteristic is often used to distribute work amongst a group of workers to parallelize the CPU and I /O.
*/
func main() {
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{10, 20, 30, 40, 50}

	var wg sync.WaitGroup

	//It receives a "receive-only" directional channel
	ch1 := generator(data1...)
	ch2 := generator(data2...)
	wg.Add(2)

	//We will loop through both the channels till all data is sent and marked as close
	go func() {
		for val := range ch1 {
			fmt.Printf("Channel1 data: %v\n", val)
		}
		wg.Done()
	}()

	go func() {
		for val := range ch2 {
			fmt.Printf("Channel2 data: %v\n", val)
		}
		wg.Done()
	}()

	wg.Wait() //Wait till all the goroutines are marked as done.
}

func generator(nums ...int) <-chan int {
	myChannel := make(chan int)

	go func() {
		//Iterate the nums data and sends it to channel.
		for _, v := range nums {
			myChannel <- v
		}
		close(myChannel)
	}()

	return myChannel
}

/*
Channel2 data: 10
Channel2 data: 20
Channel2 data: 30
Channel2 data: 40
Channel2 data: 50
Channel1 data: 1
Channel1 data: 2
Channel1 data: 3
Channel1 data: 4
Channel1 data: 5
*/
