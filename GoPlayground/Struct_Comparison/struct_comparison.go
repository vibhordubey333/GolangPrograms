package main

import (
	"fmt"
	"reflect"
)

type SampleStruct struct {
	intA     int
	stringB  string
	intSlice []int
}

func main() {

	intAObj := SampleStruct{1, "ShriRam", []int{1, 2, 3}}
	intAObj2 := SampleStruct{1, "ShriRam", []int{1, 2, 3}}
	intBObj3 := SampleStruct{2, "Jai Shri Ram", []int{3, 2, 1}}
	fmt.Println("Struct intAObj and intAObj2 are equal:", reflect.DeepEqual(intAObj, intAObj2))
	fmt.Println("Struct intAObj and intBObj3 are equal: ", reflect.DeepEqual(intAObj, intBObj3))
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 4, 3, 2}
	fmt.Println("Slice comparison:", reflect.DeepEqual(slice1, slice2))
	/*if slice1 == slice2 {
		fmt.Println("Slices are equal")
	}*/
}
