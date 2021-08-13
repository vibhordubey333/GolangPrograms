package main

import (
	"fmt"
		
)

func main() {
	
	chInt := make(chan int)
	//boolChan := make(chan bool)
	
	go SendNumbers(chInt)//,boolChan)
	
	func(){
		for v := range chInt{
			fmt.Println("Recieved : ",v)
		}
	}();
	
	//<-boolChan 		
}

func SendNumbers(chInt chan int){//,boolChan chan bool){
	for i := 0 ; i < 10 ; i++{
		fmt.Println("Sent : ",i)
		chInt <- i
	}
	close(chInt)
	//boolChan <- true
}
