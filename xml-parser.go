package main

import (
	"encoding/xml"
	"fmt"
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

func main() {
	// init struct
	var s SitemapIndex
	var n News
	newMap := make(map[string]NewsMap)
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Println("there is an error with the request firewall or something else")
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")

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

		for index, data := range newMap {
			fmt.Println("\n\n\n", index)
			fmt.Println("\n", data.Keyword)
			fmt.Println("\n", data.Location)
		}

	} else {
		fmt.Println("Argh! Broken")
	}

}
