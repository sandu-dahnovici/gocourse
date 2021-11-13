package main

import "fmt"

func main() {
	birthYear := 2003
	fmt.Println("Anii in care am trait") // using for {}
	for {
		fmt.Println(birthYear)
		if birthYear > 2021 {
			break
		}
		birthYear++
	}
}
