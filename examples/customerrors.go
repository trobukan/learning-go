package examples

import (
	"errors"
	"fmt"
)

// it's possible to use custom type as errors by implementing Error() method on them
// Here's a variant on the example above that uses a custom type to explicity represent
// an argument error.

// a custom error type usually has the suffix "Error"
type argError struct {
	arg     int
	message string
}

// adding this Error method makes argError implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// return our error
func fa(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// errors.As is a more advanced version of errors.Is.
// it checks that a given error (or any error in its chain) matches a specific 
// error type and converts to a value of that type, returning true. if there's no
// match, it returns false.
func CustomErrors() {
	fmt.Println("- - - - - Custom Errors - - - - -")
	_, err := fa(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
