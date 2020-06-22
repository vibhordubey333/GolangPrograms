package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template // To make templates visible at package level
//Template is like a container.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	tpl, err := template.ParseGlob("templates/*") // Program here will pick any file from templates dir and parse it .

	log.Println(tpl)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil) // It will print the output of the file say "I am doing great"
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
