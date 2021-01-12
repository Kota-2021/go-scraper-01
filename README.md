# go-scraper-01

This code is the my first scraper.

Access the 'sitepoint' community

git 

Get 30 article titles & links


*****

↓ Golang official web site
https://golang.org/

*****

↓about net/http
https://golang.org/pkg/net/http/

*****
used "goquery"
https://github.com/PuerkitoBio/goquery

↓about goquery
https://godoc.org/github.com/PuerkitoBio/goquery

*****

↓about HTTP request headers

request.Header.Set("cache-control", "no-cache")
→ Is the cached data the lastest ? Check this !

request.Header.Set("pragma", "no-cache")
Is it older than HTTP/1.1 ? Then look here . ↑
Same meaning as ' "cache-control", "no-cache" '

request.Header.Set("dnt", "1")
Please don't chase me .

request.Header.Set("upgrade-insecure-requests", "1")
This browser can use https. → Use the https instead of http.

request.Header.Set("referer", refererURL)
I came to this site from 'refererURL' site .

*****
