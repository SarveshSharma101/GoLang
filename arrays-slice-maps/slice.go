package main

import (
	"fmt"
	"reflect"
)

func Slice() {
	// Declaring a slice
	// 1. a := []int{}

	a := []int{}
	fmt.Println("Zero value of Slice is nil, Slice a:", a)
	fmt.Println("Length of Slice a:", len(a))
	fmt.Println("Capacity of Slice a:", cap(a))
	a = append(a, 1)
	fmt.Println("Slice a:", a)
	fmt.Println("Length of Slice a:", len(a))
	fmt.Println("Capacity of Slice a:", cap(a))
	a[0] = 2
	fmt.Println("Slice a:", a)
	fmt.Println("Length of Slice a:", len(a))
	fmt.Println("Capacity of Slice a:", cap(a))

	//2. using make
	b := make([]int, 3, 7)
	fmt.Println("Array b: ", b)

	// type of array
	fmt.Println("Slice is reference type, a reference is send to function")
	fmt.Printf("Type: %v and Kind: %v\n", reflect.TypeOf(b), reflect.TypeOf(b).Kind())

	// slice from an array

	var arr = [5]int{11, 12, 13, 14, 15}
	s := arr[1:4]
	fmt.Println("slice s: ", s)
	fmt.Printf("Type: %v and Kind: %v\n", reflect.TypeOf(s), reflect.TypeOf(s).Kind())

}
