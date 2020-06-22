package main

import (
	"log"
	"os"
	"text/template"
)

type father struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp.gohtml"))
}
func main() {
	RAM := father{
		Name:  "RAM",
		Motto: "RAM is the only truth.",
	}
	err := tpl.Execute(os.Stdout, RAM)
	if err != nil {
		log.Fatalln("Error in execute")
	}
}
