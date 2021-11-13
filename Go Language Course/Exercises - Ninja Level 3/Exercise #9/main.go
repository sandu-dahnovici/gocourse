package main

import "fmt"

func main() {
	favSport := "Football"
	switch favSport {
	case "Golf":
		fmt.Println("Let's play golf")
	case "Football":
		fmt.Println("Let's play football")
	case "Basketball":
		fmt.Println("Let's play basketball")
	default:
		fmt.Println("I think we can't play together")
	}
}
