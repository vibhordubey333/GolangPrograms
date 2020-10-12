package main

import (
	"RPC_POC/common"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	log.Println("Client started:")
	//Get RPC client by dialing at rpc.DefaultRPCPath endpoint.
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	//Create John variable of type `common.Student`
	var john common.Student

	//Get Student by ID `1`
	if err := client.Call("College.Get", 1, &john); err != nil {
		fmt.Println("Error : 1 College.Get()", err)
	} else {
		fmt.Println("Success:1 Student `%s` found with ID=1 \n ", john.FullName())
	}

	if err := client.Call("College.Add", common.Student{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}, &john); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Println("Success:2 Student `%s` created with id=1 \n")
	}
}
