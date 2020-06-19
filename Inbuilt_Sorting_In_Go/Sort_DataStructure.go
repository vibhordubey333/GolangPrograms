package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByAge []Person

/*
	We've created a slice above as indexing cannot be performed on struct i.e. we cannot do func(a Person) then index over the items.
*/

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	familyObj := []Person{
		{"A", 23},
		{"B", 21},
		{"C", 12},
	}
	sort.Sort(ByAge(familyObj))
	fmt.Println(familyObj)

}
