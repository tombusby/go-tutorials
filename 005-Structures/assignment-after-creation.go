package main

import ("fmt")

type Saiyan struct {
    Name string
    Power int
}

func main() {
    goku := Saiyan{Name: "Goku"}
    goku.Power = 9000
    fmt.Printf("%v\n", goku)

    null_sayan := Saiyan{}
    fmt.Printf("%v\n", null_sayan)
    null_sayan.Name = "The Saiyan with no name"
    fmt.Printf("%v\n", null_sayan)
}
