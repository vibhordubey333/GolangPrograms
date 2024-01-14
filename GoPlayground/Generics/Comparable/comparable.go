package main

import "fmt"

/*
In Go, the comparable constraint is used to specify that a type can be compared using the standard comparison operators (==, <, <=, >, >=).
This is an important concept in Go because some types, such as slices, maps, and functions, are not comparable.
*/

func F[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(F(1, 2))     // false
	fmt.Println(F("a", "a")) // true
	//fmt.Println(F([]int{1, 2, 3}, []int{1, 2, 3})) // Error: []int does not satisfy comparable)
}
