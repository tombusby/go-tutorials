package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 3, 4, 5}
	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)
}

// handy function for removing an item from an *unordered* slice
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	// return the slice up to (not including) the final element
	return source[:lastIndex]
}
