package main

import "fmt"

func foo() int {
	return 42
}
func bar() (int, string) {
	return 14, "Hello"
}
func main() {
	f := foo()
	b, bs := bar()
	fmt.Println("foo : ", f)
	fmt.Println("bar : ", b, bs)
}
