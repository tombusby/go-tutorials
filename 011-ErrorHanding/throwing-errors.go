package main

import (
	"errors"
	"fmt"
)

// We aren't using it here but this is the error interface
type error interface {
	Error() string
}

func main() {
	err := process(0)
	fmt.Println(err)
}

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}
