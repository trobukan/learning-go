package examples

import "fmt"

func For() {
	fmt.Println("- - - - - For - - - - -")

	fmt.Println("- - - - -")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("- - - - -")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println("- - - - -")
	for i := range 3 {
		fmt.Println("range", i)
	}
	fmt.Println("- - - - -")
	for {
		fmt.Println("Loop")
		break
	}

	fmt.Println("- - - - -")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
