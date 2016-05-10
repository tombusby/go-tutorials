package main

import (
	"fmt"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}

	// * precedence: *gohan.Father == (*gohan).Father
	fmt.Printf("%v %v\n", *gohan, *gohan.Father)
}
