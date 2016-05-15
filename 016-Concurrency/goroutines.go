package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	// The reason we need this is to stop the program terminating
	// before the go-routine can execute. This is a bad way to
	// solve that problem don't do this.
	time.Sleep(time.Millisecond * 10)
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}
