package examples

import "fmt"

// The type [n]T is an array of n values of type T
func Arrays() {
	fmt.Println("- - - - - Arrays - - - - -")
	var a[2] string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}