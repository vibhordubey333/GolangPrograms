package main

import (
	"fmt"
	"sync"
)

/*
The bridge-channel pattern refers to connect or combine multiple channels into a single channel. It allows you to merge the streams of data from
different channels into a unified stream, making it easier to consume and process the combined data.In this case, the bridge channel acts as a connector
between multiple channels, enabling the flow of data between them.
*/
func bridge(inputs ...<-chan int) <-chan int {
	output := make(chan int, len(inputs))

	go func() {
		defer close(output)

		var wg sync.WaitGroup
		wg.Add(len(inputs))

		for _, input := range inputs {
			go func(ch <-chan int) {
				defer wg.Done()
				for value := range ch {
					output <- value
				}
			}(input)
		}

		wg.Wait()
	}()

	return output
}

func main() {
	sensorData1 := make(chan int)
	sensorData2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Send values to input1 and input2 channels
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			sensorData1 <- i
		}
		close(sensorData1)
	}()
	go func() {
		defer wg.Done()
		for i := 6; i <= 10; i++ {
			sensorData2 <- i
		}
		close(sensorData2)
	}()

	// Bridge the input channels into a single output channel
	output := bridge(sensorData1, sensorData2)

	// Consume values from the output channel
	go func() {
		defer wg.Done()
		for value := range output {
			fmt.Println("Received Full Sensor Data: ", value)
		}
	}()

	wg.Wait()
}

/*
Received Full Sensor Data:  6
Received Full Sensor Data:  7
Received Full Sensor Data:  8
Received Full Sensor Data:  9
Received Full Sensor Data:  1
Received Full Sensor Data:  10
*/
