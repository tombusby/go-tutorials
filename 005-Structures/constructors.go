package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

// There are no constructors so you have a function like a factory
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	goku := NewSaiyan("Goku", 9000)
	fmt.Printf("%v\n", *goku)
}
