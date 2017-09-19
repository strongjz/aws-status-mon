package rss

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//BASEUrl - Base url to grab the AWS status page
var BASEUrl = "https://status.aws.amazon.com"

//Rss - Holds the config for the RSS and slice of Feeds
type Rss struct {
	Config *viper.Viper
	Feed   []*Feed
}

//Feed - Struct to contain the feed data for individual services
type Feed struct {
	Region  string
	Service string
	URL     string
	PollInt int
	Alert   bool
}

//PrintFeed - prints out the feed
func PrintFeed(f *Feed) string {
	return fmt.Sprintf("Service - %s : Region %s : URL %s : Poll Interval %d", f.Service, f.Region, f.URL, f.PollInt)
}

func NewRss() *Rss {

	r := Rss{}
	var f []*Feed

	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	r.Config = config
	r.Feed = f

	return &r
}

//NewFeed - Feed constructor
func NewFeed() *Feed {
	returnFeed := &Feed{}
	returnFeed.Region = "us-east-1"
	returnFeed.Service = "elasticcloudcompute"
	returnFeed.URL = fmt.Sprintf("%s/rss/%s-%s.rss", BASEUrl, returnFeed.Service, returnFeed.Region)
	returnFeed.PollInt = 60
	returnFeed.Alert = false
	return returnFeed
}

//GetFeed - grabs all the RSS feeds from the status page
func (r *Rss) GetFeed() error {

	// request and parse the front page
	resp, err := http.Get(BASEUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
		return err
	}

	feedList := parseHTML(doc)

	r.Feed = feedList

	return nil
}

func parseHTML(doc *html.Node) []*Feed {

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

	return feedList
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
	returnFeed.URL = fmt.Sprintf("%s/rss/%s-%s.rss", BASEUrl, returnFeed.Service, returnFeed.Region)
	returnFeed.PollInt = 60

	//log.Printf("Service - %s : Region %s", returnFeed.Service, returnFeed.Region)

	return returnFeed
}
