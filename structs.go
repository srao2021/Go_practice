package main

import "fmt"

type person struct {
    name string
    age  int
}

func employee(name string) *person {

    p := person{name: name}
    p.age = 25 
    return &p
}

func main() {

    fmt.Println(person{"A", 25})

    fmt.Println(person{name: "B", age: 25})

    fmt.Println(person{name: "C"})

    fmt.Println(&person{name: "D", age: 25})

    fmt.Println(employee("J"))

    s := person{name: "S", age: 25}
    fmt.Println(s.name)


}
