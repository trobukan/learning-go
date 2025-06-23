package examples

import "fmt"

// | Recursion |
// This function call itself until it reaches the base case of fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// | Closures |
// This function intSeq returns another function, which define
// anonymously in the body of intSeq. The returned functions closes
// over the variable i to form closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// | Variadic functions |
// Here's function that will take an arbitary number of ints
// as arguments.

// Within the function, the type of nums is equivalent to []int.
// we can call len(nums), iterate over it with range, etc.

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

// | Multiple Return Values |
// the (int, int) in this function signature shows that function
// that returns 2 ints
func vals() (int, int) {
	return 3, 7
}

// | Functions |
// Here's a function that takes two ints and returns their sums
// as an int
func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func Functions() {
	fmt.Println("- - - - - Functions - - - - -")
	res := plus(1, 2)
	fmt.Println(res)
	res = plusPlus(1, 2, 3)
	fmt.Println(res)

	// here we use 2 different return values from the call
	// with multiple assignment
	a, b := vals()
	fmt.Println(a, b)

	// if you only want a subject of the returned values, use
	// the blank identifier.
	_, c := vals()
	fmt.Println(c)

	// variadic functions can be called in the usual way
	// with individuals arguments
	sum(1, 2)

	// if you have already multiple args in a slice, apply them to a
	// variadic function using func(slice...) like this.
	nums := []int{29, 49, 1}
	sum(nums...)

	// we call intSeq, assigning the result (a function) to
	// nextInt. This function value captures its own i value, which
	// will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(fact(5))

	// Anonymous functions can also be recursive, but this requires explicty
	// declaring variable with var to store the function before it's defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// Since fib was previously declared in main, Go
	// knowns with function to call with fib here
	fmt.Println(fib(7))
}
