package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:Twitter`
}

func StructuredData() {

	jsonFile, err := os.Open("JsonData/UsersData.json")
	defer jsonFile.Close()
	if err != nil {
		log.Println("Error while opening file")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Error while reading file.")
	}

	var usersObj Users
	json.Unmarshal(byteValue, &usersObj)
	log.Println("Sample: ", usersObj)
	for i := range usersObj.Users {
		log.Println("UsersType:", usersObj.Users[i].Type)
		log.Println("UsersAge:", usersObj.Users[i].Age)
		log.Println("UsersName:", usersObj.Users[i].Name)
		log.Println("Facebook URL:", usersObj.Users[i].Social)
	}

}
