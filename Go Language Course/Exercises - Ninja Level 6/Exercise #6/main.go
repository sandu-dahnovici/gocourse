package main

import "fmt"

func main() {
	// anonymous func
	func() {
		fmt.Println("Hello World!")
	}()
}
