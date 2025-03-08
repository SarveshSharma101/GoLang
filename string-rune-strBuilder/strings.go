package main

import (
	"fmt"
	"strings"
)

func Strings() {
	var s string = "s"
	s = "Hello"
	fmt.Println("string s: ", s)
	fmt.Println("Indexing string")
	fmt.Println("index 0, indexing an array give the byte value: ", s[0])
	fmt.Println("to get string value of indexed value: ", string(s[0]))

	fmt.Println("Slicing a sstring: ", s[1:3])

	str := "world"

	fmt.Println("concatiniating string ", s+str)

	fmt.Println("strings are slice of byte: ", []byte(s))

	b := []byte(s)

	fmt.Println("byte to string: ", string(b))

	// string functions

	fmt.Println("to upper ", strings.ToUpper(s))
	fmt.Println("to lower ", strings.ToLower(s))
	fmt.Println("has prefix ", strings.HasPrefix(s, "he"))
	fmt.Println("has suffix ", strings.HasSuffix(s, "llo"))
	fmt.Println("Replace ", strings.Replace(s, "l", "e", 2))
	fmt.Println("Index ", strings.Index(s, "l"))
}
