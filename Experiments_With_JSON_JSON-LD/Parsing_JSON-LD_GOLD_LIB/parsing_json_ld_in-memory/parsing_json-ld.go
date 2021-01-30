package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/piprate/json-gold/ld"
)

const (
	LINE_BREAK = "************************************************************************************"
)

func main() {

	//	text := `{"@context": ["http://schema.org", { "image": { "@id": "schema:image", "@type": "@id"}  }],"id": "http://www.wikidata.org/entity/Q76","type": "Person","name": "Barack Obama","givenName": "Barack","familyName": "Obama","jobTitle": "44th President of the United States","image": "https://commons.wikimedia.org/wiki/File:President_Barack_Obama.jpg"}`
	/*text := `{
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
	}`*/
	text := `{
		"@context": "http://www.w3.org/ns/odrl.jsonld",
    	"@type": "Set",
		"ID":"http://example.com/policy:9001",
		"Assigner":"http://example.com/sony:10",
		"Assignee":"http://example.com/billie:888",
		"Name":"TestData",
		"ExpiryData":"2012-10-31 15:50:13.793654 +0000 UTC",
		"Rules":[
		   {
			  "target":"http://example.com/music:4545",
			  "type":"string",
			  "action":"http://www.w3.org/ns/odrl/2/copy",
			  "constraints":[
				 {
					"name":"http://www.w3.org/ns/odrl/2/count",
					"operator":"http://www.w3.org/ns/odrl/2/lteq",
					"rightoperand":"1"
				 }
			  ]
		   }
		],
		"Constraints":[
		   {
			  "operator":"http://www.w3.org/ns/odrl/2/eq",
			  "leftoperand":"0.50",
			  "rightoperand":"http://www.w3.org/2001/XMLSchema#decimal"
		   }
		]
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
