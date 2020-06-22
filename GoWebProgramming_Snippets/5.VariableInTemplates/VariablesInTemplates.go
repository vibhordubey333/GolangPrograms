package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/temp1.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "temp1.gohtml", "Shaken not stirred.")
	if err != nil {
		log.Fatalln("Error while executing template:", err)
	}
}
