package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Gdata saves the URL and text retrieved by scraping.
type Gdata struct {

	// url saves the URL retrieved by scraping.
	url string

	// cont saves the text retrieved by scraping.
	cont string
}

// getListing does scraping from listingURL.
// Width returns the value of the URL and text.
func getListing(listingURL string, refererURL string) ([]Gdata, int) {

	// getdata saves the value retrieved by scraping as  Gdata structure.
	var getdata []Gdata

	// count save the number of acquisitions retrieved by scraping.
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

	// 200 is success status
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

	// This code will scrape from this url site
	// referrer is the site of the site before it came url site.
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
