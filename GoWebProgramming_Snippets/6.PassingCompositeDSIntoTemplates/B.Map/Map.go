package main

import (
	"log"
	"os"
	"text/template"
)

var tpl1 *template.Template
var tpl2 *template.Template

func init() {
	tpl1 = template.Must(template.ParseFiles("../templates/temp.gohtml"))
	tpl2 = template.Must(template.ParseFiles("temp2.gohtml"))
}

func main() {
	eatable := map[string]string{
		"India":      "Rice",
		"USA":        "Corn",
		"Antarctica": "Ice",
	}
	err := tpl1.Execute(os.Stdout, eatable)
	if err != nil {
		log.Fatalf("Error while executing: ", err)
	}
	err = tpl2.Execute(os.Stdout, eatable)
	if err != nil {
		log.Fatalf("Error while executing: ", err)
	}
}
