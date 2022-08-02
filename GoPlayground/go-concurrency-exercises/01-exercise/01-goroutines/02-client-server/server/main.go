package main

import (
	"io"
	"log"
	"net"
	"time"
)

var id int64

func main() {
	//write server program to handle concurrent client connections.
	listenerObject, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal("Error encountered on server: ", err)
	}
	for {
		conn, err := listenerObject.Accept()
		if err != nil {
			continue //If error is nil listen to next connection.
		}
		go handleConn(conn) //Calling handleConn concurrently so it can handle multiple clients
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "Response from server. Received request \n ")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
