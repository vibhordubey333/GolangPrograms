package main

import "fmt"

type A struct {
	value1 int
	value2 int
}

func (a *A) Add(v *A) { //Performing addition of two objects.
	a.value1 += v.value1
	a.value2 += v.value2
}

func main() {
	x, y := &A{1, 1}, &A{2, 2}
	x.Add(y)
	fmt.Println(x)

	a, b := A{4, 1}, A{1, 4}
	a.Add(&b)
	fmt.Println(a)
}
