package main

import (
"fmt"
"time"
)

func some_function(n int) {
for i :=0; i<n; i++ {
fmt.Println(i)
}
}

func main() {
go some_function(10)
time.Sleep(time.Second)
}
