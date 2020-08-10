package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func UnStructuredData() {

	openResponse, err := os.Open("JsonData/UsersData.json")
	defer openResponse.Close()
	if err != nil {
		log.Println("Error while opening file: ", err)
	}

	byteResponse, err := ioutil.ReadAll(openResponse)
	if err != nil {
		log.Println("Error while reading file.", err)
	}

	var usersObj interface{}
	err = json.Unmarshal(byteResponse, &usersObj)
	if err != nil {
		log.Println("Error while unmarshalling data.", err)
	}
	log.Println("Sample: ", usersObj)
	log.Println("%T:", usersObj)

}
