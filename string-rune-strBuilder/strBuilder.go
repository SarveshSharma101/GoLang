package main

import (
	"fmt"
	"strings"
)

func StringBuilder() {
	var sb strings.Builder
	fmt.Println("String builder ", sb)
	fmt.Println("Capacity ", sb.Cap())
	fmt.Println("Length ", sb.Len())
	fmt.Println("Grow 10 ")
	sb.Grow(10)
	fmt.Println("Capacity ", sb.Cap())
	fmt.Println("Length ", sb.Len())

	i, err := sb.WriteString("hello")
	if err != nil {
		fmt.Println("Error writing string ", err)
	}
	fmt.Println("len of string written ", i)
	fmt.Println("sb ", sb)
	i, err = sb.Write([]byte("World"))
	if err != nil {
		fmt.Println("Error writing string ", err)
	}
	fmt.Println("len of string written ", i)
	fmt.Println("sb ", sb.String())

	fmt.Println("reset sb")
	sb.Reset()
	fmt.Println("sb ", sb)
}
