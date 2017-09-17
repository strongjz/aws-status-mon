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

func init() {
	var ErrorMessages []string

	ErrorMessages = append(ErrorMessages, "Increased Error Rates")
	ErrorMessages = append(ErrorMessages, "Intermittent API Latency")
}

func PollFeed(f []*Feed) {

	log.Printf("Starting Polling")
	log.Printf("Polling %d Services", len(f))

	numGoroutines := len(f)

	for i := 0; i < 10; i++ {

		log.Printf("Polling Starting %s-%s", f[i].Service, f[i].Region)

		go poll(f[i])

		for diff := range goroutineDelta {
			numGoroutines += diff
			if numGoroutines == 0 {
				log.Printf("Polling Finished")
				os.Exit(0)
			}
		}

	}
}

func poll(f *Feed) {

	log.Printf("Polling %s in %s", f.Service, f.Region)

	for {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(f.URL)

		findError(feed, f.Service, f.Region)

		//log.Printf("All good in the hood")
		time.Sleep(10 * time.Second)

		//goroutineDelta <- +1

	}

}

func findError(parsedFeed *gofeed.Feed, service, region string) {

	//nothing := false
	//whatsup := "nothing"

	if len(parsedFeed.Items) > 0 {
		for _, i := range parsedFeed.Items {
			if strings.Contains(i.Title, "Increased Error Rates") {
				//			nothing = true
				//			whatsup = i.Description

				alert.StandardOut(service, region, i.Description)

			}
		}
	} else {
		if strings.Contains(parsedFeed.Title, "Increased Error Rates") {
			//		nothing = true
			//		whatsup = parsedFeed.Description

			alert.StandardOut(service, region, parsedFeed.Description)
		}
	}

	log.Printf("All good in Service %s - Region %s", service, region)
	return //whatsup, nothing
}