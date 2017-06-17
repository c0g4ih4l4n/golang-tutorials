package methods

import (
	"fmt"
	"time"
)

// MyError :
// type error for demo
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v %v", e.When, e.What)
}

func run() *MyError {
	return &MyError{
		time.Now(),
		"Something wrong was happen",
	}
}

// PrintDemoError :
// Demo for interface error
func PrintDemoError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
