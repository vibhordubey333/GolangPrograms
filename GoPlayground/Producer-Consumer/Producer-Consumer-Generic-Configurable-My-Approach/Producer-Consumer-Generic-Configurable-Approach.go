// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	// It supports various types of datatypes
	data []interface{} = []interface{}{"It's a general syntax", struct{}{}, 3343.2, "test1", "test2", 344, 343, 223}
)

type Config struct {
	producerCount int
	consumerCount int
}

func main() {
	//Configurable 
	configObject := &Config{
		producerCount: 1,
		consumerCount: 10,
	}

	ch := make(chan interface{})

	wg.Add(configObject.producerCount)
	wg.Add(configObject.consumerCount)

	for i := 0; i < configObject.producerCount; i++ {
		go producer(ch)
	}

	for i := 0; i < configObject.consumerCount; i++ {
		go consumer(ch)
	}

	wg.Wait()
}

func producer(ch chan interface{}) {
	defer close(ch)
	//dataObject := new(Data)
	for _, value := range data {
		fmt.Println("Producing item : ", value)
		ch <- value
	}
	wg.Done()
}

func consumer(ch chan interface{}) {
	for v := range ch {
		switch v.(type) {
		case string:
			fmt.Println("Consuming String Type. Item: ", v)
		case struct{}:
			fmt.Println("Consuming Struct Type. Item: ", v)
		case float64, float32:
			fmt.Println("Consuming Float Type. Item: ", v)
		case int, int64, int32:
			fmt.Println("Consuming Int Type. Item: ", v)
		}
	}
	wg.Done()
}
/*
-------Output-----------
Producing item :  It's a general syntax
Producing item :  {}
Consuming String Type. Item:  It's a general syntax
Consuming Struct Type. Item:  {}
Producing item :  3343.2
Producing item :  test1
Consuming Float Type. Item:  3343.2
Consuming String Type. Item:  test1
Producing item :  test2
Producing item :  344
Consuming String Type. Item:  test2
Producing item :  343
Producing item :  223
Consuming Int Type. Item:  344
Consuming Int Type. Item:  223
Consuming Int Type. Item:  343

Program exited.
*/
