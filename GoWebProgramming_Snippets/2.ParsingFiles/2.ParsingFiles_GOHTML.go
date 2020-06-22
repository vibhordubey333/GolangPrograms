package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
		Purpose of writing this program to achieve , parsing a gohtml file and display it's content on console  or write in  file.
	*/
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("Error while parsing files : ", err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error while creating file :", err)
	}

	//err = tpl.Execute(os.Stdout, nil) // It will print content of tpl.gohtml on console.
	err = tpl.Execute(nf, nil) // It will write content tpl.gohtml in file.
	if err != nil {
		log.Fatalln("Error while executing :", err)
	}
}
