package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	// Note added & before data type to assign the pointer to the var
	goku := &Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	Super(goku)
	// power is changed due to passing (a copy of) the pointer instead
	// of a copy of the struct. I also dereference the pointer below.
	fmt.Printf("%v\n", *goku)

	// Go does not have the . vs -> distinction for getting members of
	// data values or pointed-to values.
	fmt.Printf("%v\n", goku.Power)

	// We can also print the value of the pointer itself
	fmt.Printf("%v\n", &goku)
}

// Note added * before data type to denote a pointer (like C/C++)
func Super(s *Saiyan) {
	s.Power += 10000
}
