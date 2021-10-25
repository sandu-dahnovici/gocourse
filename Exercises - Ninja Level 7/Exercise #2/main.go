package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) changeMe() {
	p.name = "Ion"
	p.age = 32
}

func main() {
	x := person{
		name: "Sandu",
		age:  17,
	}
	fmt.Println(x)
	x.changeMe()
	fmt.Println(x)
}
