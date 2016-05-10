package main

import ("fmt")

type Saiyan struct {
    Name string
    Power int
}

func main() {
    // The examples below are equivalent

    goku := &Saiyan{
        Name: "Goku",
        Power: 9000,
    }

    goku2 := new(Saiyan)
    goku2.Name = "Goku"
    goku2.Power = 9000

    fmt.Printf("%v %v\n", goku, goku2)
}
