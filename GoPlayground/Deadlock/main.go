package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 2)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	fmt.Println("Reading: 1 ", <-ch)
	fmt.Println("Reading: 2 ", <-ch)
}
