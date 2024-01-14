package main

import "fmt"

type DT interface {
	int | float64 | float32 | int8 | string
}

func Add[T DT](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(3, 4))     // Output: 7
	fmt.Println(Add(3.1, 4.1)) // Output: 7.199999999999999
	fmt.Println(Add("Hello", "World"))
}
