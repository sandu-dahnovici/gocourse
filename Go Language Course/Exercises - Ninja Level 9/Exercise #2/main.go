package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Person ", p.first, " speaks")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Sandu"}
	p.speak()
	saySomething(&p)
}
