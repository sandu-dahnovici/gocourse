package main

import (
	// va fi link github "dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	c1 := canine{
		name: "Spike",
		age:  5,
	}
	// dog.Years(c1.age)
	fmt.Println(c1)

}
