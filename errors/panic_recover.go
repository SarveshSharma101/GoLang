package main

import "fmt"

func F1(recover bool) {
	if recover {
		defer recoverPanic()
	}
	defer fmt.Println("Defer before panic, panic from f1")
	F2()
	fmt.Println("After panic F1")
}
func F2() {
	defer fmt.Println("Panic from f2")
	panic("panic")
	fmt.Println("After panic from f2")
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovering")
	}
}
