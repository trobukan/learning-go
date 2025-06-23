package examples

import "fmt"

func Maps() {
	fmt.Println("- - - - - Maps - - - - -")
	var a = map[string]string{
		"brand": "Ford", 
		"model": "Mustang",
		"year": "1964"}
	// OR
	var c = make(map[string]string) // empty
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	b := map[string]int{
		"Oslo": 1,
		"Bergen": 2,
		"Trondheim": 3,
		"Stavanger": 4}

	a["color"] = "red"
	delete(a, "year")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a["brand"])
	
	// Check for specific elements in map
	val1, ok1 := a["brand"]
	_, ok2 := a["day"]
	fmt.Println(val1, ok1)
	fmt.Println(ok2)

	// Map are references
	d := b
	d["Oslo"] = 6
	fmt.Println(d, b)

	// Iterate over maps
	numbers := map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
	}
	var order []string // defining the order
	order = append(order, []string{"one", "two", "three", "four"}...)

	// Looping with no order
	for k, v := range numbers {
		fmt.Println(k, v)
	}
	// Looping with the defined order
	for _,element := range order {
		fmt.Println(element, numbers[element])
	}

}

