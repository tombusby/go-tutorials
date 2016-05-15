package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		// Once all the workers are sleeping, the code calling the channel
		// blocks until a worker becomes available.
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	// The select statement can be used to determine if the
	// channel accepted the data
	for {
		select {
		case c <- rand.Int():
		default:
			// this can be left empty to silently drop the data
			fmt.Println("dropped")
		}

		time.Sleep(time.Millisecond * 50)
	}
}
