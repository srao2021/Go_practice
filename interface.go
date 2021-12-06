package main

import (
	"fmt"
)

type people interface {
	name() string
}

type employee struct {
	age float64
}

func (e employee) name() string {
	return "abc"
}

func give_details(p people) {
	fmt.Println(p)
	fmt.Println(p.name())
}

func main() {
	p := employee{age: 35}
	give_details(p)
}
