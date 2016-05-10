package main

import ("fmt")

type Saiyan struct {
    Name string
    Power int
}

func main() {
    goku := Saiyan{
        Name: "Goku",
        Power: 9000,
    }
    Super(goku)
    // goku's power is unchanged due to a copy being passed to Super()
    fmt.Printf("%v\n", goku)
}

func Super(s Saiyan) {
    s.Power += 10000
}
