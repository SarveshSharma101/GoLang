package main

import (
	"errors"
	"fmt"
)

func main() {
	Errors()
	WrapUnWrapError()
	PanicRecover()
}
func Errors() {
	fmt.Println("error using ftm.Errorf")
	fmt.Println(fmt.Errorf("this ftm error"))

	fmt.Println("error using errors.new")
	fmt.Println(errors.New("this is new error"))
}

func WrapUnWrapError() {
	ce1 := ReturnCustomError()
	e1 := fmt.Errorf("Wrapping c1: %w", ce1)
	fmt.Println("Wrapping error")
	fmt.Println("-->", e1)
	fmt.Println("Wrapping error")
	fmt.Println("-->", errors.Unwrap(e1))

	fmt.Println("comparing 2 errors")
	fmt.Println("compare using == ")
	fmt.Println("-->", e1 == ce1)
	fmt.Println("Using is")
	fmt.Println("-->", errors.Is(e1, ce1))
}

type CustomError struct {
	Msg string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("error:%v ", c.Msg)
}

func ReturnCustomError() error {
	return CustomError{
		Msg: "this is custom error",
	}
}

func PanicRecover() {
	fmt.Println("calling panic function")
	deferWithPanic()
	defer fmt.Println("panic in parent")
}

func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovering from panic: ", r)
	}
}

func deferWithPanic() {
	defer RecoverFromPanic()

	panic("panicing")
}
