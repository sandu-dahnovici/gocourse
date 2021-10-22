package main

import "fmt"

func main() {
	s := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."}}
	for _, v := range s {
		for _, w := range v {
			fmt.Println(w)
		}
		fmt.Println()
	}
}
