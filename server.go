package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello word</h1>")
}
func NewsAggPageHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "hello  test titre", News: "la belle vie du bon vivant"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", NewsAggPageHandler)
	http.ListenAndServe(":8080", nil)
}
