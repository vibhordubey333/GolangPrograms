// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var structObject struct{}
	var interfaceObject interface{}
	assertType("String Type")
	assertType(structObject)
	assertType(interfaceObject)
}

func assertType(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("Received String Type")
	case struct{}:
		fmt.Println("Received Struct Type")
	case interface{}:
		fmt.Println("Received Empty-Interface Type")
	default:
		fmt.Println("Unknown Type")
	}
}

/*
Output:
Received String Type
Received Struct Type
Unknown Type
 */
