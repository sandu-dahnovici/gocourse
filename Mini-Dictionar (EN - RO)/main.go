package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]string)
	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content := bufio.NewScanner(file)
	words_file := []string{}
	for content.Scan() {
		words_file = append(words_file, content.Text())
	}
	for i, v := range words_file {
		if i%2 == 0 {
			words[v] = words_file[i+1]
		}
	}
	var v string
	for v != "0" {
		fmt.Println("1 : Adaugare de la tastatura")
		fmt.Println("2 : Stegerea unui cuvant")
		fmt.Println("3 : Afisarea traducerea unui cuvant englez")
		fmt.Println("4 : Afisarea tuturor cuvintelor")
		fmt.Println("5 : Salvarea dictionarului in fisier")
		fmt.Println("0 : Iesire")
		fmt.Scanln(&v)
		switch v {
		case "1":
			var en, ro string
			fmt.Print("Engleza : ")
			fmt.Scanln(&en)
			fmt.Print("Romana : ")
			fmt.Scanln(&ro)
			words[en] = ro
		case "2":
			var en string
			fmt.Print("Introduceti cuvantul in engleza care doriti sa il stergeti : ")
			fmt.Scanln(&en)
			if _, ok := words[en]; !ok {
				fmt.Println("Acest cuvant nu exista")
			} else {
				delete(words, en)
				fmt.Println("Sters cu succes!")
			}
		case "3":
			var en string
			fmt.Print("Introduceti cuvantul de cautat : ")
			fmt.Scanln(&en)
			if ro, ok := words[en]; !ok {
				fmt.Println("Acest cuvant nu exista")
			} else {
				fmt.Printf("%s : %s\n", en, ro)
			}
		case "4":
			for en, ro := range words {
				fmt.Printf("%s : %s\n", en, ro)
			}
		case "5":
			f, err := os.OpenFile("words.txt", os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			for en, ro := range words {
				f.WriteString(en)
				f.WriteString("\n")
				f.WriteString(ro)
				f.WriteString("\n")
			}
			fmt.Println("Salvat!")
		default:
			fmt.Println("Ati introdus varianta gresita!")
		}
	}
}
