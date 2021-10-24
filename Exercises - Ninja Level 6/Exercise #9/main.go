package main

import "fmt"

func sqr(x int) int {
	return x * x
}

func sqr4(sqr func(x int) int, y int) int {
	return sqr(y) * sqr(y)
}

func main() {
	// x^4
	fmt.Println(sqr4(sqr, 2))
}
