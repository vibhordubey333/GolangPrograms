package main

/*
A. Empty struct struct{} is realized in a special way in Go.
B.  It’s a smallest building block in Go. It’s size is literally 0 bytes.
C. If it has zero size. you may create a slice of 1000’s empty structures and this slice will be very tiny. Because really Go stores only a number of them in 
the slice but not them itself. The same story with channels.
D. Very useful in channels when you notify about some event but you don’t need to pass any information about it, only a fact. Best solution is to pass an empty 
structure because it will only increment a counter in the channel but not assign memory, copy elements and so on. Sometime people use Boolean values for this purpose, 
but it’s much worse.
E. Zero size container for methods. You may want have a mock for testing interfaces. Often you don’t need data on it, just methods with predefined input and 		output.
F. Go has no Set object. Bit can be easily realized as a map[keyType]struct{}. This way map keeps only keys and no values.
*/
//I usually use it where I would have used a channel of booleans. ie, instead of;

/*
//****** Without using empty struct{} 
func main() {
    done := make(chan bool, 1)

    go func() {
        // simulate long running task
        time.Sleep(4 * time.Second)
        done <- true
        fmt.Println("long running task is done")
    }()

    <-done
    close(done)

    fmt.Printf("whole program is done.")
}
*/

import (
	"fmt"
	"time"
)

func main() {
	structCh := make(chan struct{})
	fmt.Println("Task Launched:")
	go printHelloWorld(structCh)

	<-structCh
	fmt.Println("Main GoRoutine Is Closing.")
}

func printHelloWorld(structCh chan struct{}) {
	fmt.Println("In Progress Please Wait.")
	time.Sleep(5 * time.Second)
	structCh <- struct{}{}
}
