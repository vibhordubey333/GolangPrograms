package main

type Num interface {
	~int | ~float32 | ~float64 | ~string | ~bool
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

/*
If we don't use tilde operator ?
Then, error will be encountered.
Integer does not satisfy Num (possibly missing ~ for int in Num)
Why did we encounter this error? After all, Integer is a type of int, and the Num type set contains int. So, why do we get this error?

The answer lies in the fact that the type parameter of the Compare function is bound to the Num type set. Consequently, it only allows types that are
explicitly defined within the type set. However, you might argue that since Integer has an underlying type of int, it should be allowed. This is where the
underlying operator (~) comes into play.

To accept any type as an argument whose underlying type matches the set of types defined within the type set, we need to use the ~ operator. It allows
us to explicitly look for underlying types that correspond to the types defined in the type set.

TL;DR
It allows us to explicitly look for underlying types that correspond to the types defined in the type set.
*/
