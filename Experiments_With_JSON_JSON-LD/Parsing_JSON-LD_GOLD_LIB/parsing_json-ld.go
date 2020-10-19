package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	jsonld "github.com/emersion/go-jsonld"
	"github.com/piprate/json-gold/ld"
)

const (
	LINE_BREAK = "************************************************************************************"
)

// ODRL OptiSAM-NG schema
type person struct {
	ID    string           `jsonld:"@id"`
	Name  string           `jsonld:"name"`
	URL   *jsonld.Resource `jsonld:"url"`
	Image *jsonld.Resource `jsonld:"image"`
}
type policy struct {
	ID         string `jsonld:"@id"`
	Assigner   string
	Assignee   string
	Name       string    `jsonld:"name"`
	ExpiryDate time.Time `jsonld:"name"`
	Rules      []Rules   `jsonld:rules`
	Constraint []Constraints

	//	URL   *jsonld.Resource `jsonld:"url"`
	//	Image *jsonld.Resource `jsonld:"image"`
}
type Rules struct {
	Target     string
	Type       string //Permission/Prohibition
	Constraint []Constraints
	Action     []Actions
}

type Constraints struct {
	LeftOperand  string
	RightOperand string //Confirm from ODRL
	Operator     string
}

type Actions struct {
	Actions    string
	Constraint Constraints
}

func main() {

	//text := `{"@context": ["http://schema.org", { "image": { "@id": "schema:image", "@type": "@id"}  }],"id": "http://www.wikidata.org/entity/Q76","type": "Person","name": "Barack Obama","givenName": "Barack","familyName": "Obama","jobTitle": "44th President of the United States","image": "https://commons.wikimedia.org/wiki/File:President_Barack_Obama.jpg"}`
	text := `{
		"@context": "http://www.w3.org/ns/odrl.jsonld",
		"@type": "Set",
		"uid": "http://example.org/policy/5",
		"permission": [{
			"target": "http://example.com/asset/1",
			"assignee": "http://example.com/party/Alice",
			"action": "use"
		}],
		"prohibition": [{
			"target": "http://example.com/asset/1",
			"assignee": "http://example.com/party/Bob",
			"action": "use"
		}]
	}`
	var result map[string]interface{}
	json.Unmarshal([]byte(text), &result)

	fmt.Println("\n\nSimple Representation.")
	fmt.Println("\n", LINE_BREAK+"\n")
	for k, v := range result {
		fmt.Println("K: ", k, " :", v)
	}
	fmt.Println("\n", LINE_BREAK+"\n")
	fmt.Println("Result : ", result)
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	expandResult, err := proc.Expand(result, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}
	ld.PrintDocument("JSON-LD expansion succeeded", expandResult)

}
