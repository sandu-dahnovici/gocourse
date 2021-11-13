package main

import "fmt"

func main() {
	a := 76
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x", b, b, b)
}
