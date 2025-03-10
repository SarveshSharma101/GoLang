package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(GetCustomError())
	e := GetCustomError()
	fmt.Println(WrapCustomErr(e))
	//wrapping error
	we := WrapCustomErr(e)
	fmt.Println("Wrapped error: ", we)
	fmt.Printf("Unwrapping:%v\n ", errors.Unwrap(we))

	//checking wrapped error
	fmt.Println("using = ", e == we)
	fmt.Println("using errors.Is ", errors.Is(we, e))

	//panic
	F1(true)
}
