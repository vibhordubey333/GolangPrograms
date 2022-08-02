package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	//This will be called when there are no instances available in the pool.
	New: func() interface{} { //Below is the anonymous functions that'll return interface{}
		fmt.Println("Allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}

func log(w io.Writer, debug string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")
	w.Write(b.Bytes())
	//After using placing pool object back to the pool.
	bufPool.Put(b)
}
