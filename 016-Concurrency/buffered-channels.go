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
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// We allow the channel to buffer up to 100 items. However this doesn't
	// really solve the underlying problem.
	c := make(chan int, 100)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	for {
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}
