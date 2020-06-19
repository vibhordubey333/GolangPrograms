package main

import (
	"fmt"
	"sort"
)

/*type family struct {
	Name string
	Age  int
}*/

func main() {
	familyObj := []struct {
		Name string
		Age  int
	}{
		{"A", 23},
		{"B", 1},
		{"C", 1},
		{"E", 5},
		{"F", 7},
		{"G", 90},
		{"H", 1},
	}

	sort.SliceStable(familyObj, func(i, j int) bool {
		if i < j {
			return familyObj[i].Age < familyObj[j].Age
		} else if i == j {
			// If age is equal then sort on the basis of name.
			return familyObj[i].Name < familyObj[j].Name
		} else {
			return familyObj[i].Age < familyObj[j].Age
		}
		return true
	})

	fmt.Println(familyObj)
}
