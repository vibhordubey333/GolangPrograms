package main

import (
	//"encoding/json"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello ProtoBuf")

	vibhor := &Person{
		Name: "Vibhor",
		Age:  "19",
		SocialFollowers: &SocialFollowers{
			Youtube: 999,
			Twitter: 100,
		},
	}
	data, err := proto.Marshal(vibhor)
	//Marshalling is used to convert struct to bytes.
	if err != nil {
		log.Fatal("Marshalling error.")
	}
	fmt.Println("Data :", data) // It will print data in ASCII format.
	// Unmarshalling is used to convert bytes to struct
	newVibhor := &Person{}
	err = proto.Unmarshal(data, newVibhor)
	if err != nil {
		log.Fatalf("Unmarshalling error")
	}

	fmt.Println(newVibhor.GetAge())
	fmt.Println(newVibhor.GetName())
	fmt.Println(newVibhor.SocialFollowers.GetTwitter())
	fmt.Println(newVibhor.SocialFollowers.GetYoutube())

}
