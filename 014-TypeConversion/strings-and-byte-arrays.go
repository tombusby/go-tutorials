package main

import "fmt"

func main() {
	// Strings are immutable, so we end up with three copies
	str_a := "The spice must flow"
	bytes := []byte(str_a)
	str_b := string(bytes)

	fmt.Println(str_a)
	fmt.Println(bytes)
	fmt.Println(str_b)
}
