package main

import "fmt"

// ByteCounter type
type ByteCounter int
//If we remove below implementation then we'll get error in main method as ByteCounter type doesn't implement Write() method.
func (b *ByteCounter) Write(p []byte) (n int, err error){
	*b += ByteCounter(len(p))
	return len(p),nil
}


// Implement Write method for ByteCounter
// to count the number of bytes written.

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
//Output
// 11