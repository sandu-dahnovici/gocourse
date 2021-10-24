package main

import "fmt"

func foo() {
	fmt.Println("Defer here")
}

func main() {
	defer foo()
	fmt.Println("Hello there")
}
