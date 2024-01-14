package main

type Num interface {
	int | float32 | float64 | string | bool
}

func Compare[T Num](a, b T) bool {
	return a == b
}

func main() {

	type Integer int

	var a Integer = 1
	var b Integer = 2

	println(Compare(a, b))
}
