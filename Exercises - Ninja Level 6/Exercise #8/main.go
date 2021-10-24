package main

import "fmt"

func sqr() func() int {
	return func() int {
		return 42
	}
}
func main() {
	f := sqr()
	fmt.Println(f())
}
