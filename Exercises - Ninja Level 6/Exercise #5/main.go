package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	l float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.l * s.l
}

func (c circle) area() float64 {
	const pi = 3.14
	return pi * 2 * c.r
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{4}
	c := circle{6.5}
	info(s)
	info(c)
}
