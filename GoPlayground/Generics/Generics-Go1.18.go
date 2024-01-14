package main

import (
    "fmt"
)

type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
    ~float32 | ~float64
}

type Number interface {
    Integer | Float
}

func Min[T Number](x, y T) T {
    if x < y {
        return x
    }
    return y
}

type IntType int64

func main() {
    x, y := 1, 3
    a, b := 1.1, 4.5
    k, v := IntType(-12344), IntType(12)
    fmt.Println(Min(x, y))
    fmt.Println(Min(a, b))
    fmt.Println(Min(k, v))
}
