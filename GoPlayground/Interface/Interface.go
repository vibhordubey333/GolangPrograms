package main

import "fmt"

type Humans interface {
	Speak() string
	//NotImplemented() // Methods which are not implemented will give errors. Uncomment then see.
}

type Indian struct{}

func (i Indian) Speak() string {
	return "Hindi, " + "Tamil," + " Kannad"
}

type Russian struct{}

func (r Russian) Speak() string {
	return "Russian"
}

func main() {
	//Method 1 to access using interface object.
	hObject := []Humans{Indian{}, Russian{}}

	for index, object := range hObject {
		fmt.Println("i:", index, " object:", object.Speak())
	}
	fmt.Println("---------------------------------------------------------")
	//Method 2 to access using interface object.
	var iObject Humans = Indian{}
	fmt.Println("Indian Struct:", iObject.Speak())

}
