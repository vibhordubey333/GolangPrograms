package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		//In case args is not provided use the default one.
		name = "I am champion."
	}

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	</head>
	`
	//fmt.Println(tpl)
	//Instead of printing content on the console , printing in terminal.
	nf, err := os.Create("index.html")
	defer nf.Close()
	if err != nil {
		log.Fatal("Error while creating file : ", err)
	}
	io.Copy(nf, strings.NewReader(tpl))

}
