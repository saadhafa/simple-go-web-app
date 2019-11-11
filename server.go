package main

import ("fmt"
				"net/http")

func index_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `
		<h1>hello word <h2>
	`)
}

func main(){
	http.HandleFunc("/",index_page)
	http.ListenAndServe(":8080", nil)
}