package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintln(x, "\t", y, "\t", z)
	fmt.Println(s)
}
