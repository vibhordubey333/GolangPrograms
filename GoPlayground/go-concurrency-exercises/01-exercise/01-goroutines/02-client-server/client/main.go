package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//connect to server on localhost port 9999
	connectionObject, err := net.Dial("tcp", "localhost:9999")
	defer connectionObject.Close()
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, connectionObject) // To display server message on console screen.
}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
