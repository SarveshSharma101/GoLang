package main

import (
	"fmt"
	"reflect"
)

func Arrays() {
	// Declaring an array
	// 1. var a [5]int
	var a [5]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	fmt.Println("Array a: ", a)

	//2. var b = [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b: ", b)

	//3. c :=[...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array c: ", c)

	// type of array
	fmt.Println("Array is value type, a copy of array is send to function")
	fmt.Printf("Type: %v and Kind: %v\n", reflect.TypeOf(a), reflect.TypeOf(a).Kind())
}
