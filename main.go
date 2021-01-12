package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Gdata struct {
	url  string
	cont string
}

func getListing(listingURL string, refererURL string) ([]Gdata, int) {
	var getdata []Gdata
	var count int

	//HTTP client with timeout
	client := &http.Client{
		//timeout 30 second
		Timeout: 30 * time.Second,
	}
	// create request
	request, err := http.NewRequest("GET", listingURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	//Setting headers
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("dnt", "1")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", refererURL)
	// do request with '↑ request headers'
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// 200 is success
	if resp.StatusCode == 200 {
		//get body data
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//get a data in body
		docf := doc.Find(".link-top-line a")
		//number of searches
		count = docf.Length()
		//take out one by one from docf
		docf.Each(func(i int, s *goquery.Selection) {
			//get href data
			link, _ := s.Attr("href")
			// Gdata is struct
			var gd Gdata
			// Put in Gdata
			gd.url = link
			gd.cont = s.Text()
			//Gdata put in Array
			getdata = append(getdata, gd)
		})
	}
	return getdata, count
}
func main() {

	fmt.Println("started")

	set := map[string]string{
		"url":     "https://www.sitepoint.com/community/",
		"referer": "https://www.sitepoint.com/",
	}
	fmt.Println("set : ", set)

	link, count := getListing(set["url"], set["referer"])

	fmt.Println("count : ", count)

	for _, s := range link {
		fmt.Println(s.url)
		fmt.Println(s.cont)
	}

	fmt.Println("finished")
}
