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
/*
Output:
List:  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100]
Total Sum:  5050
*/
