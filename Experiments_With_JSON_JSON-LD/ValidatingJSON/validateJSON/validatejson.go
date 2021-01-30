package main

import (
	"log"

	schema "github.com/lestrrat-go/jsschema"
	"github.com/lestrrat-go/jsval/builder"
)

//For validating JSON as well as JSON-LD schema.
func main() {

	
	s, err := schema.ReadFile("D:\\Golang\\GolangPrograms\\Experiments_With_JSON_JSON-LD\\JSON_Test_Data\\json-ld_parsing.json")
	if err != nil {
		log.Printf("failed to open schema: %s", err)
		return
	}
	b := builder.New()
	v, err := b.Build(s)
	if err != nil {
		log.Printf("failed to build validator: %s", err)
		return
	}
	var input interface{}
	if err := v.Validate(input); err != nil {
		log.Printf("validation failed: %s", err)
		return
	}
}
