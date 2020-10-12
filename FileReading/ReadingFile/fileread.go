package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("D:\\Golang\\GolangPrograms\\FileReading\\FilesToRead\\json-ld_parsing_2.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	printData(string(data))
}
func printData(data interface{}) {
	fmt.Println("Data: ", data)
}
