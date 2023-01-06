package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Problem Statement:
Compute sum from 1 to 100 i.e 5050

Constraints:
1. Compute sum using 3 goroutines.
2. Cannot define length explicitly each goroutine will pick.
------------------------------------------------------------------------
Solution:
We can solve this problem using producer-consumer pattern.
1. Define a pool of integers say arr.
2. Define 1 producer and 3 consumers.
3. Producer will read from the pool of integers i.e. arr
4. And Consumer will keep on listening and pass further to computeSum which will keep on adding sum atomically.
5.
*/

/*
 */
var (
	wgConsumer   sync.WaitGroup
	wgProducer   sync.WaitGroup
	wgComputeSum sync.WaitGroup
	totalSum     uint64
)

type SumList struct {
	value []int
}

type Config struct {
	producerCount int
	consumerCount int
}

func main() {
	ch := make(chan int)

	valueObject := new(SumList)
	valueObject.value = make([]int, 0)
	for i := 1; i <= 100; i++ {
		valueObject.value = append(valueObject.value, i)
	}
	fmt.Println("List: ", valueObject.value)
	configObject := new(Config)
	configObject.producerCount = 1
	configObject.consumerCount = 3

	wgProducer.Add(configObject.producerCount)
	wgConsumer.Add(configObject.consumerCount)

	for i := 0; i < configObject.producerCount; i++ {
		go producer(ch, valueObject)
	}

	for j := 0; j < configObject.consumerCount; j++ {
		go consumer(ch)
	}

	wgProducer.Wait()
	wgConsumer.Wait()
	fmt.Println("Total Sum: ", totalSum)
}

func producer(ch chan int, sumObject *SumList) {
	defer close(ch)
	for i := 0; i < len(sumObject.value); i++ {
		ch <- sumObject.value[i]
	}
	wgProducer.Done()
}

func consumer(ch chan int) {

	for v := range ch {
		wgComputeSum.Add(1)
		go computeSum(v)
	}
	wgComputeSum.Wait()
	wgConsumer.Done()
}

func computeSum(ch int) {
	atomic.AddUint64(&totalSum, uint64(ch))
	wgComputeSum.Done()
}
