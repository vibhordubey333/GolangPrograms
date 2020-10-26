package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/piprate/json-gold/ld"
)

type Policy struct {
	Context     string        `jsonld:"context"`
	Type        string        `jsonld:"type"`
	ID          string        `jsonld:"@id"`
	Assigner    string        `jsonld:"assigner"`
	Assignee    string        `jsonld:"assignee"`
	Name        string        `jsonld:"name"`
	ExpiryDate  time.Time     `jsonld:"name"`
	Rules       []Rules       `jsonld:"rules"`
	Constraints []Constraints `jsonld:"constraints"`

	//	URL   *jsonld.Resource `jsonld:"url"`
	//	Image *jsonld.Resource `jsonld:"image"`
}
type Rules struct {
	Target     string        `jsonld:"target"`
	Type       string        `jsonld:"type"` //Permission/Prohibition
	Constraint []Constraints `jsonld:"constraint"`
	Actions    []Actions     `jsonld:"actions"`
}

type Constraints struct {
	LeftOperand  string `jsonld:"leftoperand"`
	RightOperand string `jsonld:"rightoperand"` //Confirm from ODRL
	Operator     string `jsonld:"operator"`
}

type Actions struct {
	Actions    string      `jsonld:"actions"`
	Constraint Constraints `jsonld:"constraint"`
}

func main() {

	/*text := `{
		"ID": "http://example.com/policy:9001",
		"Assigner": "http://example.com/sony:10",
		"Assignee": "http://example.com/billie:888",
		"Name": "TestData",
		"ExpiryData": "2012-10-31 15:50:13.793654 +0000 UTC",
		"Rules": [{
			"target": "http://example.com/music:4545",
			"action": "http://www.w3.org/ns/odrl/2/copy",
			"constraints": [{
					"name": "http://www.w3.org/ns/odrl/2/count",
					"operator": "http://www.w3.org/ns/odrl/2/lteq",
					"rightoperand": "1"
			}],
		}],
		"Constraints": [{
			"name": "http://www.w3.org/ns/odrl/2/payAmount",
			"operator": "http://www.w3.org/ns/odrl/2/eq",
			"rightoperand": "0.50",
			"rightoperanddatatype": "http://www.w3.org/2001/XMLSchema#decimal",
			"rightoperandunit": "http://cvx.iptc.org/iso4217a/AUD"
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

	var result Policy
	json.Unmarshal([]byte(text), &result)
	fmt.Println("Result : ", result)
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	marshallObject, err := json.Marshal(result)
	fmt.Println("Marshall Object :", string(marshallObject))

	//Storing []byte data into map.
	mapObject := make(map[string]interface{})
	errMap := json.Unmarshal(marshallObject, &mapObject)
	errHandler(errMap)

	expandResult, err := proc.Expand(mapObject, options)
	if err != nil {
		log.Println("Error when expanding JSON-LD document:", err)
		return
	}
	ld.PrintDocument("JSON-LD expansion succeeded", expandResult)

	/*
		var policy Policy
		r := bytes.NewReader([]byte(text))
		policyErr := json.NewDecoder(r).Decode(&policy)
		errHandler(policyErr)
		fmt.Println("Decoded output: ", policy)

		fmt.Println("Decoded Array : ")
		for k, v := range policy.Rules {
			fmt.Println("K: ", k, " ||  V: ", v)
		}
		//var PolicyMapObject
		marhsalObj, err := json.Marshal(policy)
		errHandler(err)
		fmt.Println("Marshal Object : ", string(marhsalObj))

		//Storing []byte data into map.
		mapObject := make(map[string]interface{})
		errMap := json.Unmarshal(marhsalObj, &mapObject)
		errHandler(errMap)

		proc := ld.NewJsonLdProcessor()
		options := ld.NewJsonLdOptions("")

		expandResult, err := proc.Expand(mapObject, options)
		if err != nil {
			log.Println("Error when expanding JSON-LD document:", err)
			return
		}
		ld.PrintDocument("JSON-LD expansion succeeded", expandResult)
	*/
}
func errHandler(err error) {
	if err != nil {
		fmt.Println("Error :", err.Error)
		return
	}
}
