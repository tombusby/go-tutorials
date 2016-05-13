package main

import "fmt"

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["Krillen"] = &Saiyan{
		Name:    "Krillen",
		Friends: make(map[string]*Saiyan),
	}

	fmt.Println(goku)
}
