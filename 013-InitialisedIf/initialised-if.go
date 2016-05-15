package main

import (
	"errors"
	"fmt"
)

func main() {
	// err will be out of scope outside the if/else/elseif blocks
	if err := process(); err != nil {
		// We can use this to handle errors passed back
		fmt.Println(err)
	}
}

func process() error {
	return errors.New("Some random error")
}
