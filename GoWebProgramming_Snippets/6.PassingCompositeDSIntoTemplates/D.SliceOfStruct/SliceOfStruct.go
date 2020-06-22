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
	Krishna := father{
		Name:  "Krishna",
		Motto: "RAM and Krishna both are ultimate and same.",
	}
	Temp := father{
		Name:  "Dummy",
		Motto: "This is test",
	}
	fatherSliceOfStruct := []father{RAM, Krishna, Temp}
	err := tpl.Execute(os.Stdout, fatherSliceOfStruct)
	if err != nil {
		log.Fatalln("Error in execute")
	}
}
