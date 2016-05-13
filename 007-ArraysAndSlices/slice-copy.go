package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)

	// Copy too few values: the remainder are zeroed
	worst2 := make([]int, 5)
	copy(worst2, scores[:2])
	fmt.Println(worst2)

	// Copy too many values: no error, only first N copied
	worst3 := make([]int, 5)
	copy(worst3, scores[:5])
	fmt.Println(worst3)

	// Essentially the same as a "copy too many" but with an offset
	worst4 := make([]int, 5)
	copy(worst4[2:4], scores[:5])
	fmt.Println(worst4)
}
