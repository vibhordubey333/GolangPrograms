package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := os.Open("notavailable.txt")
	fmt.Println("Error Message: ", err)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file does not exist.")
		} else if errors.Is(err, io.ErrUnexpectedEOF) {
			fmt.Println("Unexpected EOF encountered.")
		} else {
			fmt.Println("An unknown error occurred:", err)
		}
	}
}

/*
Output:
Error Message:  open notavailable.txt: no such file or directory
The file does not exist.
*/
