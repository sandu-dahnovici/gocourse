package main

import "fmt"

type person struct {
	first string
	last  string
	ficf  []string
}

func main() {
	p1 := person{
		first: "Sandu",
		last:  "Dahnovici",
		ficf:  []string{"Strawberry", "Cherry", "Banana"},
	}
	p2 := person{
		first: "Catalina",
		last:  "Dahnovici",
		ficf:  []string{"Vanilla", "Apple", "Pineapple"},
	}
	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, Person := range m {
		fmt.Println("Nume : ", Person.first)
		fmt.Println("Prenume : ", Person.last)
		fmt.Print("Arome favorite de inghetata : ")
		for _, v := range Person.ficf {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
