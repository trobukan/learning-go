package examples

import (
	"fmt"
	"strconv"
)

func TypeCasting() {
	fmt.Println("- - - - - Type Casting - - - - -")
	// Basic type casting
	// Anywhere we need to change the type of the variable, the typecasting is needed.
	var a int = 32 
	f := float64(a)
	fmt.Println(f) 

	// Conversions between string and int
	var s string = "42"
	v, _ := strconv.Atoi(s)
	fmt.Println(v)

	var i int = 42
	s = strconv.Itoa(i)
	fmt.Println(s)

	// Conversions between int and float
	f = 12.29420
	i = int(f)
	fmt.Println(i)

	i = 34
	f = float64(i)
	fmt.Println(f)

	// Strings and bytes
	s = "Hello World"
	var b []byte = []byte(s) // convert to bytes

	fmt.Println(b)

	ss := string(b) // convert to string
	fmt.Println(ss)

	// Type conversion during division
	a = 6/3
	f = 6.3/3

	fmt.Println(a, f)
}
