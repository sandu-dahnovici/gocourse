package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

var items = map[string]int{
	"guitar":   500,
	"drums":    1500,
	"keyboard": 1000,
}

func init() {
	tpl = template.Must(tpl.ParseGlob("templates/[a-z]*.html"))
}

func main() {
	http.Handle("/main/", http.StripPrefix("/main/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send_email)
	fmt.Println("The server started listeting at 8080...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/main/", http.StatusMovedPermanently)
}

func send_email(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	username := req.FormValue("username")
	quantity, _ := strconv.Atoi(req.FormValue("quantity"))
	selected_item := req.FormValue("items")
	price_item := items[selected_item]
	order_price := quantity * price_item
	data := struct {
		Username      string
		Quantity      int
		Selected_item string
		Price_item    int
		Order_price   int
	}{username, quantity, selected_item, price_item, order_price}
	err = tpl.ExecuteTemplate(w, "send.html", data)
	if err != nil {
		panic(err)
	}
}
