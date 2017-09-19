package rss

import (
	"log"

	gofeed "github.com/mmcdole/gofeed"
	"github.com/strongjz/aws-status-mon/alert"
	"os"
	"strings"
	"time"
)

var goroutineDelta = make(chan int)

//ErrorMessages - List of strings to indicate an error in the rss
var errorMessages []string

func init() {

	errorMessages = append(errorMessages, "Increased Error Rates")
	errorMessages = append(errorMessages, "Intermittent API Latency")
}

//PollFeed - Starts polling the list of feed sent
func (r *Rss) PollFeed() {

	//log.Printf("Starting Polling")
	//log.Printf("Polling %d Services", len(f))

	numGoroutines := len(r.Feed)

	for _, i := range r.Feed {
		log.Printf("[INF0] Polling Starting %s-%s", i.Service, i.Region)
		go poll(i)
	}

	for diff := range goroutineDelta {
		numGoroutines += diff
		if numGoroutines == 0 {
			log.Printf("Polling Finished")
			os.Exit(0)
		}
	}
}

func poll(f *Feed) {

	log.Printf("[INF0] Polling %s in %s", f.Service, f.Region)

	for {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(f.URL)

		findError(feed, f.Service, f.Region)

		time.Sleep(time.Duration(f.PollInt) * time.Minute)

	}

}

func findError(parsedFeed *gofeed.Feed, service, region string) {

	if len(parsedFeed.Items) > 0 {
		for _, i := range parsedFeed.Items {
			if statusCheck(i.Title) {

				//log.Printf("\nTitle: %s \n Description: %s\n", i.Title, i.Description)
				alert.StandardOut(service, region, i.Description)
			}
		}
	} else {
		if statusCheck(parsedFeed.Title) {
			//log.Printf("Title: %s \n Description: %s\n", parsedFeed.Title, parsedFeed.Description)
			alert.StandardOut(service, region, parsedFeed.Description)
		}
	}

	log.Printf("[INF0] All good in Service %s - Region %s", service, region)
	return
}

func statusCheck(t string) bool {

	for _, s := range errorMessages {
		if strings.Contains(t, s) && !strings.Contains(t, "RESOLVED") {
			return true

		}
	}

	return false
}
