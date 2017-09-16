package feed

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetFeed() ([]string, error) {
	// request and parse the front page
	resp, err := http.Get("https://status.aws.amazon.com/")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var feedList []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "rss") {
					feedList = append(feedList, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return feedList, nil
}
