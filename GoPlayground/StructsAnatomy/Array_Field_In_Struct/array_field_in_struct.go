package main

import (
	"fmt"
)

type SliceInStruct struct {
	Id     int
	Names  []string
	Cities []string
}

var structObject *SliceInStruct

func init() {
	cityData := []string{"KNP", "LCK", "MUM"}
	namesData := []string{"Andrew Kirby", "Alexi Murdoch", "Olafur Arnalds"}
	structObject = &SliceInStruct{}
	structObject.Cities = make([]string, 0)
	structObject.Names = make([]string, 0)
	structObject.Cities = cityData
	structObject.Names = namesData
}
func main() {
	fmt.Println("Slice in struct.")
	fmt.Println("Names:", structObject.Names)
	fmt.Println("Cities:", structObject.Cities)
}
