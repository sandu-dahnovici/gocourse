package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Home")
	})
	http.HandleFunc("/dog/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Dog page")
	})
	http.HandleFunc("/me", func(rw http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(rw, "Sandu")
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
