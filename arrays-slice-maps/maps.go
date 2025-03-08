package main

import "fmt"

func Maps() {
	// Declaring map
	var m map[int]string
	fmt.Println("uninitialized maps ", m)
	// m[1] = "a" // error
	m = map[int]string{}
	fmt.Println("initialized map ", m)
	m[1] = "a" //no error
	fmt.Println(m[1])

	// using make
	n := make(map[string]string)
	n["key"] = "value"
	fmt.Println("n: ", n)

	// check if key exist
	_, ok := n["key"]
	fmt.Println("exist? ", ok)

	//removing a value
	fmt.Println("-->", m)

	delete(m, 1)
	fmt.Println("-->", m)
}
