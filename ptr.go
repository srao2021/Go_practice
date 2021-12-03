package main

import "fmt"

func val_return(a int) {
a = 0
}

func pointer_accept(a *int) {
*a =0
}

func main() {

b :=10
fmt.Println("first: ", b)
val_return(b)
fmt.Println("second: ", b)
pointer_accept(&b)
fmt.Println("third: ", b)
}
