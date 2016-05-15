package main

import (
	"fmt"
)

type Add func(a int, b int) int

func main() {
	still_in_scope := 4

	// here we define a closure, which inherits the parent scope
	result := process(func(a int, b int) int {
		return a + b + still_in_scope
	})

	fmt.Println(result)
}

func process(adder Add) int {
	return adder(1, 2)
}
