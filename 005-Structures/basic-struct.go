package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Printf("%v\n", goku)

	// Rely on implicit ordering of the struct
	goku2 := Saiyan("Goku", 9000)
	fmt.Printf("%v\n", goku2)
}
