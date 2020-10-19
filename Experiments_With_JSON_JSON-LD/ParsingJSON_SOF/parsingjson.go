package main

import (
	"bytes"
	"fmt"
	"time"

	jsonld "github.com/emersion/go-jsonld"
	"github.com/piprate/json-gold/ld"
)

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

	text := `{"@context": ["http://schema.org", { "image": { "@id": "schema:image", "@type": "@id"}  }],"id": "http://www.wikidata.org/entity/Q76","type": "Person","name": "Barack Obama","givenName": "Barack","familyName": "Obama","jobTitle": "44th President of the United States","image": "https://commons.wikimedia.org/wiki/File:President_Barack_Obama.jpg"}`
	textBytes := []byte(text)
	var container person
	dec := jsonld.NewDecoder(bytes.NewReader(textBytes))
	fmt.Println("Dec : ", dec.FetchContext)
	dec.FetchContext = func(url string) (*jsonld.Context, error) {
		var fetchedContext jsonld.Context //TODO fetch the context
		return &fetchedContext, nil
	}

	err := dec.Decode(&container)
	fmt.Println("Error while unmarshalling json-ld: ", err)
	fmt.Println("Output: ", container)

	proc := ld.NewJsonLdProcessor()
	fmt.Println("PROC: ", proc)

}
