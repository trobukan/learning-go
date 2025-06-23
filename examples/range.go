package examples

import "fmt"

func Range() {
	fmt.Println("- - - - - Range - - - - -")
	nums := []int{2, 3, 6}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	// range on map iterates over key/values pairs
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	// range can also over just the keys of a map
	for k:= range kvs {
		fmt.Println(k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the
	// rune	and the second the rune itself.
	for i, c := range "go"{
		fmt.Println(i, c)
	}

}