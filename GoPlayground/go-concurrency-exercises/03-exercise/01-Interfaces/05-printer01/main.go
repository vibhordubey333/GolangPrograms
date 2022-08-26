package main

import (
	"io"
	"os"
)

func main() {
	printer(os.Stdout, "hello")
}
//For more explanation see program 05-printer
func printer(w io.Writer, str string) {
	w.Write([]byte(str))
	w.Close()
}
