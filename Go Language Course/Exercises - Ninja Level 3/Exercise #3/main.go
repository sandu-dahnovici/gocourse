package main

import "fmt"

func main() {
	birthYear := 2003
	fmt.Println("Anii in care am trait")
	for birthYear <= 2021 {
		fmt.Println(birthYear)
		birthYear++
	}
}
