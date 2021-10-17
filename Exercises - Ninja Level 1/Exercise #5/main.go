package main

import "fmt"

type integer int

var x integer
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 15
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
