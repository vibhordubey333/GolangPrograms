package main

import (
	"flag"
	"fmt"
)

// Reference: https://gobyexample.com/command-line-flags
func main() {
	//Usage Syntax: String(name string,value string,usage string
	wordObject := flag.String("Args1", "Value1", "String Type")

	//Important To Call : After all flags are defined, call flag.Parse
	flag.Parse()

	fmt.Println("WordObject: ", *wordObject)
	/*
		Output With Args:
		 ./Flag.exe -Args1=fe
		WordObject:  fe
	*/

	/*
		Output Without Args:
		$ ./Flag.exe
		WordObject:  Value1
	*/

}
