package main

import "fmt"

func main() {
	// Anonymous struct
	messi := struct {
		first_last_name string
		matches         int
		goals           int
	}{
		first_last_name: "Lionel Mesii",
		matches:         975,
		goals:           643,
	}
	fmt.Println("First and last name : ", messi.first_last_name)
	fmt.Println("Matches : ", messi.matches)
	fmt.Println("Goals : ", messi.goals)
}
