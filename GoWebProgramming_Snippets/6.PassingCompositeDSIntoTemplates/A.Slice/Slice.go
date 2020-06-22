package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var tp2 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../templates/temp.gohtml"))
	tp2 = template.Must(template.ParseFiles("../templates/temp2.gohtml"))
}
func main() {
	sages := []string{"Srila Parubhupada", "Chaitanya Mahaprabhu", "Goswami Tulsidas"}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln("Error while executing.")
	}
	log.Println("______________________*****************************_______________________")
	err = tp2.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln("Error while executing temp2.gohtml :", err)
	}
}
