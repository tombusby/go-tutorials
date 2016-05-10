package main

import (
	"fmt"
)

func getPower() int {
	return 9001
}

func main() {
	power := getPower()
	fmt.Printf("It's over %d\n", power)
}
