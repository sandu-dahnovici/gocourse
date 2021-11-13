package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println("Numarul ", i, " se imparte exact la 4")
		} else if i%2 == 0 {
			fmt.Println("Numarul ", i, " este par")
		} else {
			fmt.Println("Numarul ", i, " e impar")
		}
	}
}
