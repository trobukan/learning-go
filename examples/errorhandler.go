package examples

import (
	"errors"
	"fmt"
)

// In go it's idiomatic to communicate errors via an explicit, separate return value.
// This contracts with the expections used in languages like java and ruby and the overloaded
// single result/error value sometimes used in C. go's approach makes it easy to see
// which functions return errors and to handle them using the same language constructs
// employed to other, non-error tasks

// By convention, errors are the latest return value and type error,
// a built-in interface.
// errors.New constructs a basic error value with the given error
// message.
// A nil value in the error position indicates that there was no error.
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// A sentinel error is a predeclared variable that's used to signify
// a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

// we can wrap errors with higher-level to add context. the simplest way to do this
// is with the %w verb in fmt. Errorf. Wrapped errors create a logical chain
// (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and
// errors.As
func makeTea(arg int) error {
	switch arg {
	case 2:
		return ErrOutOfTea
	case 4:
		return fmt.Errorf("making tea: %w", ErrPower)

	}
	return nil
}

func ErrorHandler() {
	fmt.Println("- - - - - Error Handler - - - - -")

	// it's common to use an inline error check in the if line
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// errors.Is checks that a given error (or any error in its chain) matches a
	// a specific error value. This is especially useful with wrapped or nested errors,
	// allowing you to identify error types or sentinel errors in a chain errors
	for i := range 5 {
		if err := makeTea(i); err != nil {
			switch {
			case errors.Is(err, ErrOutOfTea):
				fmt.Println("we should buy more tea")
			case errors.Is(err, ErrPower):
				fmt.Println("Now it's dark")
			default:
				fmt.Printf("unknown error: %s\n", err)
			}
		}
	}
}
