package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("https://status.aws.amazon.com/")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	var feed []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "rss") {
					//fmt.Printf("https://status.aws.amazon.com%s\n", a.Val)
					feed = append(feed, a.Val)

					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	fmt.Printf("Number of Feeds %d\n", len(feed))
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	i := random.Intn(len(feed))
	fmt.Printf("Random Feed %s\n", feed[i])

}
