package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("defer 1 i is ", i)
	i = i + 5 //5
	defer fmt.Println("defer 2 i is", i)
	fmt.Println("i is ", i) //5
}

/*
Output:
i is  5
defer 2 i is 5
defer 1 i is  0
*/
