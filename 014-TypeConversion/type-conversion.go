package main

import "fmt"

func main() {
	add(1, 2)
}

// This is a pointless function, but I'm just demonstrating the use
// of the "empty interface": `interface{}`
func add(a interface{}, b interface{}) interface{} {
	// Type casting is done with .(TYPE)
	sum := a.(int) + b.(int)

	// You can also use a switch (no fall-through with go switches)
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case string, bool:
		fmt.Println("a is a bool or a string")
	default:
		fmt.Println("a is... something else")
	}

	return sum
}
