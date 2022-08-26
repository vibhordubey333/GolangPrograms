package main

import (
	"os"
)

func main() {
	printer(os.Stdout, "hello")
}
//As we can close the file also it results into a bad code as printer should only be used for printing.
//For correct way see printer01 program
func printer(f *os.File, str string) {
	f.Write([]byte(str))
	f.Close()
}

//	os.Stdout.Write([]byte("hello"))
//	os.Stdout.Close()
