package main

import ("fmt"
				"net/http"
				"io/ioutil")



func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	byt, _ := ioutil.ReadAll(resp.Body)
	string_body := string(byt)
	fmt.Println(string_body)
	resp.Body.Close()
}