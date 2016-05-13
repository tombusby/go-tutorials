package main

import (
	"fmt"
)

func main() {
	names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10)
	var names2 []string
	scores := make([]int, 0, 20)

	fmt.Printf("%v\n%v\n%v\n%v\n", names, checks, names2, scores)
}
