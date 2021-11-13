package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`}}
	for name, fav := range m {
		fmt.Print(name, "'s favourite things : ")
		for _, obj := range fav {
			fmt.Print(obj, " , ")
		}
		fmt.Println()
	}

	m[`Sandu`] = []string{"Football", "Golang", "Family"}

	for name, fav := range m {
		fmt.Print(name, "'s favourite things : ")
		for _, obj := range fav {
			fmt.Print(obj, " , ")
		}
		fmt.Println()
	}
}
