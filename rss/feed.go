package rss

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var BASEUrl = "https://status.aws.amazon.com/rss"

type Feed struct {
	Region  string
	Service string
	URL     string
	PollInt int
	Alert   bool
}

func PrintFeed(f *Feed) string {
	return fmt.Sprintf("Service - %s : Region %s : URL %s : Poll Interval %d", f.Service, f.Region, f.URL, f.PollInt)
}

func NewFeed() *Feed {
	returnFeed := &Feed{}
	returnFeed.Region = "us-east-1"
	returnFeed.Service = "elasticcloudcompute"
	returnFeed.URL = fmt.Sprintf("https://status.aws.amazon.com/rss/%s-%s.rss", returnFeed.Service, returnFeed.Region)
	returnFeed.PollInt = 60
	returnFeed.Alert = false
	return returnFeed
}

func GetFeed() ([]*Feed, error) {
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

	var feedList []*Feed

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "rss") {

					parsed := parseFeed(a.Val)
					if parsed != nil {
						feedList = append(feedList, parsed)
					}
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

func parseFeed(f string) *Feed {
	//AWS Service feed url are /rss/SERVICE-REGION.rss
	returnFeed := &Feed{}

	sep1 := strings.Index(f, "-")

	sep2 := strings.Index(f, ".")

	if sep1 == -1 || sep2 == -1 {
		return nil
	}

	returnFeed.Service = f[5:sep1]
	returnFeed.Region = f[sep1+1 : len(f)-4]
	returnFeed.URL = fmt.Sprintf("%s/%s-%s.rss", BASEUrl, returnFeed.Service, returnFeed.Region)
	returnFeed.PollInt = 60

	//log.Printf("Service - %s : Region %s", returnFeed.Service, returnFeed.Region)

	return returnFeed
}