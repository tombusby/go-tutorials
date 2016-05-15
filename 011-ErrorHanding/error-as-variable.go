package main

import (
	"fmt"
	"io"
)

func main() {
	var input int
	_, err := fmt.Scan(&input)
	// We're comparing the error returned agaisnt one in the lib
	if err == io.EOF {
		fmt.Println("no more input!")
	}
}
