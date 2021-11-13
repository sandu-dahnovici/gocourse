package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I'm %v years old\n", p.last, p.first, p.age)
}

func main() {
	me := person{
		first: "Sandu",
		last:  "Dahnovici",
		age:   17,
	}
	me.speak()
}
