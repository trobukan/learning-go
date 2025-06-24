package examples

import (
	"fmt"
)

// This person type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct
// with the given  name.

// go is a garbage collected language, you can safely return a pointer to
// a local variable - it will only be cleaned up by the garbage collector
// when there are no active references to it
func newPerson(name string) *person {
	return &person{
		name: name,
		age:  42,
	}
}

func Structs() {
	fmt.Println("- - - - - Structs - - - - -")
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	// Omit fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// it's a idiomatic to encapsulate new struct in 
	// constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // acess struct fields with a dot

	// you can also use dots with struct pointers - the pointers are 
	// automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutables
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type is only used for a single value, we don't have to give it
	// a name. the value can have an anonymous struct type.
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
