package main

import (
	"log"

	jsonld "github.com/emersion/go-jsonld"
)

type Person struct {
	ID     string   `jsonld:"@id"`
	Name   string   `jsonld:"http://schema.org/name"`
	Offers []string `jsonld:"http://schema.org/offers"`
}

/*const jsonString = `{
  "@id": "http://me.markus-lanthaler.com/",
  "http://schema.org/name": "Manu Sporny"
}`*/

const jsonString = `{

	"@context" : "http://schema.org",
	"@type": "Product",
	"name": "Moleskine Notebook",
	"offers" : [
	   {
		"@type" : "Offer",
		"description": "Sold By Amazon",
		"priceSpecification" :
			 {
			  "@type" : "PriceSpecification",
			  "currency": "GBP",
			  "value": 9.88
			 }
		}]
 }`e

/*
 const jsonString = `{

	"@context" : "http://schema.org",
	"@type": "Product",
	"name": "Moleskine Notebook",
	"offers" : [
	   {
		"@type" : "Offer",
		"description": "Sold By Amazon",
		"priceSpecification" :
			 {
			  "@type" : "PriceSpecification",
			  "currency": "GBP",
			  "value": 9.88
			 }
		}]
 }`*/

func main() {

	var p Person
	if err := jsonld.Unmarshal([]byte(jsonString), &p); err != nil {
		log.Fatal(err)
	}

	log.Println(p.ID+" "+p.Name+" ", len(p.Offers))
}

/*   // It is giving error :  jsonld: fetching remote contexts is disabled
func main() {

	text := `{"@context": ["http://schema.org", { "image": { "@id": "schema:image", "@type": "@id"}  }],"id": "http://www.wikidata.org/entity/Q76","type": "Person","name": "Barack Obama","givenName": "Barack","familyName": "Obama","jobTitle": "44th President of the United States","image": "https://commons.wikimedia.org/wiki/File:President_Barack_Obama.jpg"}`
	textBytes := []byte(text)
	var container interface{}
	err := jsonld.Unmarshal(textBytes, container)
	fmt.Println("Error while unmarshalling json-ld: ", err)
	fmt.Println("Output: ", container)
}
*/
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
