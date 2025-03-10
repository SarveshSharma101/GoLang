package main

import "fmt"

type customError struct {
	errname   string
	errstring string
}

func (c *customError) Error() string {
	return fmt.Sprintf("%v: %v", c.errname, c.errstring)
}

func GetCustomError() error {
	cr := customError{
		errname:   "new-error",
		errstring: "this is a custom error",
	}
	return &cr
}

func WrapCustomErr(e error) error {
	return fmt.Errorf("wrapping custom error: %w", e)
}
