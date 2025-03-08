package main

import "fmt"

func Runes() {
	s := "Hello world"

	r := []rune(s)
	fmt.Println(fmt.Sprintf("Rune of %v is %v", s, r))
}
