package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	// prints "0 false" - 0 is default int and key doesn't exist
	fmt.Println(power, exists)

	// prints 1 since only one entry
	total := len(lookup)
	fmt.Println(total)

	delete(lookup, "goku")
}
