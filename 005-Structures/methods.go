package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

// Note the type declaration before the function name
func (s *Saiyan) Super() {
	s.Power += 10000
}

func main() {
	goku := &Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	// function is called as a method
	goku.Super()

	// Power value has been changed by Super()
	fmt.Printf("%v\n", *goku)
}
