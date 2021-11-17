package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Home")
	})
	http.HandleFunc("/dog/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Dog page")
	})
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "It's me Sandu")
	})
	http.ListenAndServe(":8080", nil)
}
