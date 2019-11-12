package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "encoding/xml"
)

// type SitemapIndex struct{
// 	Location []Location `xml:"sitemap"`
// }

// type Location struct{
// 	Loc string `xml:"loc"`
// }

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
			fmt.Println("this time its the body err")
		}

		body_string := string(bodyByt)
		fmt.Println(body_string)
		resp.Body.Close()

	} else {
		fmt.Println("Argh! Broken")
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	//fmt.Println(body)

	// if resp == nil {
	// 	fmt.Println("there is an err")
	// }else{
	// 	byt, _ := ioutil.ReadAll(resp.Body)
	// 	fmt.Println(byt)
	// }
	// // string_body := string(byt)
	// //fmt.Println(string_body)
	// fmt.Println("all fine")
	// resp.Body.Close()

}
