package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	d := make(chan int, 2)

	go func() {
		c <- 42
	}()
	d <- 90
	fmt.Println(<-c, <-d)
}
