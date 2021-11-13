package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 10; j < 20; j++ {
				c <- j
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
