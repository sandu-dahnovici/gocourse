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
	fmt.Println("Nume : ", p1.last)
	fmt.Println("Prenume : ", p1.first)
	fmt.Print("Arome favorite de inghetata : ")
	for _, v := range p1.ficf {
		fmt.Print(v, "\t")
	}
	fmt.Println()
	fmt.Println("Nume : ", p2.last)
	fmt.Println("Prenume : ", p2.first)
	fmt.Print("Arome favorite de inghetata : ")
	for _, v := range p2.ficf {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}
