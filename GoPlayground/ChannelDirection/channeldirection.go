package main

import (
	"fmt"
)

func main() {
	receiveCh := make(chan string)
	sendCh := make(chan string)

	//Spinning goroutine
	go genMsg(receiveCh)
	go relayMsg(receiveCh, sendCh)

	//Recieve message on sendCh
	v := <-sendCh
	fmt.Println(v)
}

func relayMsg(receiveMsg <-chan string, sendMsg chan<- string) {
	//Recieve message on recive channel
	msg := <-receiveMsg
	//Send it to send channel
	sendMsg <- msg
}

func genMsg(sendChannel chan<- string) {
	//Send on channel
	sendChannel <- "message"
}
