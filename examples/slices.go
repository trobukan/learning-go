package examples

import "fmt"

func Slices() {
	fmt.Println("- - - - - Slices - - - - -")
	// A common way of declaring a slice
	myslice := []int{1, 2, 3}
	// Two functions that can return length and capacity of a slice
	fmt.Println(myslice, len(myslice), cap(myslice))

	// Create a slice from an array
	arr := [6]int{1, 2, 3, 4, 5, 6}
	myslice = arr[2:4]

	fmt.Println(myslice, len(myslice), cap(myslice))

	// Create a slice from a make() function
	myslice1 := make([]int, 5, 10)
	fmt.Println(myslice1, len(myslice1), cap(myslice1))

	myslice2 := make([]int, 5)
	fmt.Println(myslice2, len(myslice2), cap(myslice2))

	// Acess elements of a slice
	prices := []int{10, 20, 30}
	fmt.Println(prices[0], prices[1])

	// Change elements of a slice
	prices[2] = 50
	fmt.Println(prices[2])

	// Append elements

	myslice3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myslice3)

	myslice3 = append(myslice3, 7, 8)
	fmt.Println(myslice3)

	// Append one slice to another slice

	myslice4 := []int{1, 2, 3}
	myslice5 := []int{4, 5, 6}
	myslice6 := append(myslice4, myslice5...)
	fmt.Println(myslice6)

	// Change the length of a slice
	arr = [6]int{9, 10, 11, 12, 13, 14}
	myslice = arr[1:5]
	fmt.Println(myslice, len(myslice), cap(myslice))

	myslice = arr[1:3]
	fmt.Println(myslice, len(myslice), cap(myslice))

	// Memory efficiency
	// When using slices, go loads all the underlying elements into memory
	// If the array is large and you need only a few elements, it's better to copy
	// those elements using the copy() function
	// The copy() function creates a new underlying array with only the required
	// elements for the slice. This will reduce the memory used for the
	// program.

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(numbers, len(numbers), cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Println(numbersCopy, len(numbersCopy), cap(numbersCopy))
}
