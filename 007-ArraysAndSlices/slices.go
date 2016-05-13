package main

import (
	"fmt"
)

func main() {
	// Note the lack of size between []
	scores := []int{1, 4, 293, 4, 9}
	fmt.Println(scores)

	// We can use make() to create a slice too
	scores2 := make([]int, 0, 10) // length 0, capacity 10
	// scores2[6] = 9033 // this errors
	scores2 = append(scores2, 5)
	// Index 0 will contain 5
	fmt.Println(scores2)

	// Re-slice the slice
	scores3 := make([]int, 0, 10)
	scores3 = scores3[0:6]
	scores3[5] = 9033
	fmt.Println(scores3)

	// Unlike python/ruby, a slice points to the same array (not a copy)
	scores4 := []int{1, 2, 3, 4, 5}
	scores4_slice := scores4[2:4]
	scores4_slice[0] = 999
	fmt.Println(scores4)

	// Re-slice the slice
	scores5 := make([]int, 0, 10)
	scores5 = scores5[0:6]
	scores5[5] = 9033
	fmt.Println(scores5)

	// Unlike python/ruby, a slice points to the same array (not a copy)
	scores6 := []int{1, 2, 3, 4, 5}
	scores6_slice := scores6[2:4]
	scores6_slice[0] = 999
	fmt.Println(scores6)

	// Go doesn't support python-style negative numbers in slices
	scores7 := []int{1, 2, 3, 4, 5}
	scores7 = scores7[:len(scores7)-1]
	fmt.Println(scores7)

	appendDemo()

	counterIntuive()
}

// A slice can grow beyond its capacity.
// Go grows by calculating x2 on the original capacity
func appendDemo() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)
	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		if cap(scores) != c {
			old_c := c
			c = cap(scores)
			fmt.Printf("Cap: %d (old) %d (new), i: %d\n", old_c, c, i)
		}
	}
}

func counterIntuive() {
	scores := make([]int, 5)
	scores = append(scores, 9332)
	fmt.Println()
	// Output is [0 0 0 0 0 9332] since scores already contains five 0s
	fmt.Println(scores)
}
