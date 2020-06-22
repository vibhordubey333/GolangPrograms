package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}
type car struct {
	Manufacturer string
	Model        string
}
type items struct {
	Wisdom    []sage
	Transport []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp.gohtml"))
}
func main() {
	sageObj1 := sage{
		Name:  "RAM",
		Motto: "RAM is the only truth.",
	}
	sageObj2 := sage{
		Name:  "Krishna",
		Motto: "RAM and Krishna both are ultimate and same.",
	}
	sageObj3 := sage{
		Name:  "Dummy",
		Motto: "This is test",
	}
	carObj1 := car{
		Manufacturer: "A1",
		Model:        "10",
	}
	carObj2 := car{
		Manufacturer: "A2",
		Model:        "11",
	}
	carObj3 := car{
		Manufacturer: "A3",
		Model:        "12",
	}
	sageObj := []sage{sageObj1, sageObj2, sageObj3}
	carObj := []car{carObj1, carObj2, carObj3}
	data := items{
		Wisdom:    sageObj,
		Transport: carObj,
	}
	//Alternative way .
	/*data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}*/
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln("Error in execute")
	}
}
