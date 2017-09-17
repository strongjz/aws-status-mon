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
func PollFeed(f []*Feed) {

	log.Printf("Starting Polling")
	log.Printf("Polling %d Services", len(f))

	numGoroutines := len(f)

	for i := 0; i < 10; i++ {

		log.Printf("Polling Starting %s-%s", f[i].Service, f[i].Region)

		go poll(f[i])

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

	log.Printf("Polling %s in %s", f.Service, f.Region)

	for {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(f.URL)

		findError(feed, f.Service, f.Region)

		time.Sleep(10 * time.Second)

	}

}

func findError(parsedFeed *gofeed.Feed, service, region string) {

	if len(parsedFeed.Items) > 0 {
		for _, i := range parsedFeed.Items {
			if statusCheck(i.Title) {
				alert.StandardOut(service, region, i.Description)
			}
		}
	} else {
		if statusCheck(parsedFeed.Title) {
			alert.StandardOut(service, region, parsedFeed.Description)
		}
	}

	log.Printf("All good in Service %s - Region %s", service, region)
	return
}

func statusCheck(t string) bool {

	for _, s := range errorMessages {
		if strings.Contains(t, s) {
			return true

		}
	}

	return false
}
