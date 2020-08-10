package main

import (
	"fmt"
)

func main() {

	var names []string

	if names == nil {
		fmt.Println("Slice is nil.")
		names = append(names, "1")
		names = append(names, "2")
		names = append(names, "3")

		fmt.Println("Array : ", names, "  Length: ", len(names), " Capacity: ", cap(names))

		names = append(names, "1")
		names = append(names, "2")
		names = append(names, "3")

		fmt.Println("Array : ", names, "  Length: ", len(names), " Capacity: ", cap(names))
	} else {
		fmt.Println("Slice is not nil.")
	}

}
