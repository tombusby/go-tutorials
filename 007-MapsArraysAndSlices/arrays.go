package main

import (
	"fmt"
)

func main() {
	// Arrays are declared with a fixed size and type
	var scores [10]int
	scores[0] = 339

	// Fields not explicitly set default to the datatype's 0-value
	fmt.Println(scores)

	// Arrays can be initialised with values
	scores2 := [4]int{9001, 9333, 212, 33}
	fmt.Println(scores2)

	// len can be used to get the length of an array
	fmt.Println("Length:", len(scores2))

	// We can iterate over an array with range
	fmt.Println()
	for index, value := range scores2 {
		fmt.Println(index, value)
	}
}
