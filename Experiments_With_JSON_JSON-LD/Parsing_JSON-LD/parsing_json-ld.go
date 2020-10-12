package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"github.com/piprate/json-gold/ld"
// )

// func main() {

// 	data, err := ioutil.ReadFile("D:\\Golang\\GolangPrograms\\FileReading\\FilesToRead\\json-ld_parsing.json")
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}

// 	proc := ld.NewJsonLdProcessor()
// 	options := ld.NewJsonLdOptions("")

// 	// expanding remote document

// 	expanded, err := proc.Expand(string(data), options)
// 	if err != nil {
// 		log.Println("Error when expanding JSON-LD document:", err)
// 		return
// 	}

// 	ld.PrintDocument("JSON-LD expansion succeeded", expanded)

// 	// expanding in-memory document

// 	doc := map[string]interface{}{
// 		"@context":  "http://schema.org/",
// 		"@type":     "Person",
// 		"name":      "Jane Doe",
// 		"jobTitle":  "Professor",
// 		"telephone": "(425) 123-4567",
// 		"url":       "http://www.janedoe.com",
// 	}

// 	expanded, err = proc.Expand(doc, options)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ld.PrintDocument("JSON-LD expansion succeeded", expanded)

// }
