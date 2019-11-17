package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Location []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Println("there is an error with the request firewall or something else")
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")

		bodyByt, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic("err with get request")
		}

		defer resp.Body.Close()

		var s SitemapIndex

		xml.Unmarshal(bodyByt, &s)

		for _, location := range s.Location {
			fmt.Printf("\n %s", location)
		}

	} else {
		fmt.Println("Argh! Broken")
	}

}
