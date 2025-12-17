package main

/*
The error type is a built-in interface
type error interface {
    Error() string
}
Functions often return an error value, and calling code should handle errors by testing whether the error equals nil
A nil error denotes success; a non-nil error denotes failure.
*/

import (
	"fmt"
	"time"
)

type TestError struct {
	When time.Time
	What string
}

func (e *TestError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &TestError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
