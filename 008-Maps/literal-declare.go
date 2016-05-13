package main

import "fmt"

func main() {
	// We can declare maps as literals
	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	// We can do a type conversion using the range keyword
	// Order of iteration is random
	for key, value := range lookup {
		fmt.Println(key, value)
	}
}
