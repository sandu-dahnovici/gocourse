package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
