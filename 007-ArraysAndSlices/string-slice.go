package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "The spice must flow"
	// This will five 4 since it's the index in the slice (not original)
	index := strings.Index(haystack[5:], " ")
	fmt.Println(index)
}
