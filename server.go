package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type SitemapIndex struct {
	Location []string `xml:"sitemap>loc"`
}

type News struct {
	Title     []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `
	<h1>Hello word <a href="/news">News</a></h1>
	`)

}

func NewsAggPageHandler(w http.ResponseWriter, r *http.Request) {

	// struct init
	var s SitemapIndex
	var n News
	newMap := make(map[string]NewsMap)

	// Get request to washingtonpost xml page

	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")

	if err != nil {
		panic("there is an error with the request firewall or something else")
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")

		// reading Response body

		bodyByt, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			panic("err with get request")
		}

		xml.Unmarshal(bodyByt, &s)

		for _, location := range s.Location {

			tempURL := strings.Replace(location, "\n", "", -1)

			resp, _ := http.Get(tempURL)

			byts, _ := ioutil.ReadAll(resp.Body)

			xml.Unmarshal(byts, &n)

			for index, _ := range n.Keywords {
				newMap[n.Title[index]] = NewsMap{n.Keywords[index], n.Locations[index]}
			}
		}

		p := NewsAggPage{Title: "This is going to take a long time to load ...", News: newMap}

		t, _ := template.ParseFiles("index.html")

		t.Execute(w, p)
	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", NewsAggPageHandler)
	http.ListenAndServe(":8080", nil)

}
