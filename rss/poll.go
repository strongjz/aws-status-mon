package rss

import (
	"log"

	gofeed "github.com/mmcdole/gofeed"
	"github.com/strongjz/aws-status-mon/alert"
	"strings"
	"time"
)

func init() {
	var ErrorMessages []string

	ErrorMessages = append(ErrorMessages, "Increased Error Rates")
	ErrorMessages = append(ErrorMessages, "Intermittent API Latency")
}

func PollFeed(f *Feed) {

	log.Printf("Polling Feed %s-%s", f.Region, f.Service)

	for {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(f.URL)

		findError(feed, f.Region, f.Service)

		//log.Printf("All good in the hood")
		time.Sleep(10 * time.Minute)
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
