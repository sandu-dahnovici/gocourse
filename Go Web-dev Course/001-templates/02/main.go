package main

import (
	"os"
	"text/template"
)

type Hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := []Hotel{
		Hotel{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "southern",
		},
		Hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		panic(err)
	}
}
