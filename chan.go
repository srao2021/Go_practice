package main

import "fmt"

func main() {
	some_channel := make(chan int)

	go func() {
		some_channel <- 10
	}()

	numb := <-some_channel
	fmt.Println(numb)
}
