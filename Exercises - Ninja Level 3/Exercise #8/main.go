package main

import "fmt"

func main() {
	birthYear := 2003
	switch {
	case 2021-birthYear > 18:
		fmt.Println("Sunt minor")
	default:
		fmt.Println("Sunt matur")
	}
}
