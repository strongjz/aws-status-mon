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
		go r.poll(i)
		time.Sleep(time.Nanosecond)
	}

	for diff := range goroutineDelta {
		numGoroutines += diff
		if numGoroutines == 0 {
			log.Printf("Polling Finished")
			os.Exit(0)
		}
	}
}

func (r *Rss) poll(f *Feed) {

	log.Printf("[INF0] Polling %s in %s", f.Service, f.Region)

	for {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(f.URL)

		r.findError(feed, f.Service, f.Region)

		time.Sleep(time.Duration(f.PollInt) * time.Minute)

	}

}

func (r *Rss) findError(parsedFeed *gofeed.Feed, service, region string) {

	//var shouldAlert bool
	//var message string

	if len(parsedFeed.Items) > 0 {
		for _, i := range parsedFeed.Items {
			if r.statusCheck(service, region, i.Title, i.Published) {

				log.Printf("[DEBUG][ALERT] %s-%s \n Title: %s \n Description: %s\n Pub Date: %s\n", service, region, i.Title, i.Description, i.Published)
				alert.Alert(r.Config, service, region, i.Description)

			}
		}
	} else {
		if r.statusCheck(service, region, parsedFeed.Title, parsedFeed.Published) {
			log.Printf("[DEBUG][ALERT] %s-%s  \n Title: %s \n Description: %s\n Pub Date: %s\n", service, region, parsedFeed.Title, parsedFeed.Description, parsedFeed.Published)
			alert.Alert(r.Config, service, region, parsedFeed.Description)

		}
	}

	log.Printf("[INF0] All good in Service %s - Region %s", service, region)
	return
}

func (r *Rss) statusCheck(service, region, title, time string) bool {

	log.Printf("[DEBUG][STATUS CHECK] %s-%s: Title: %s Time: %s", service, region, title, time)

	status := false
	runErrorMsg := true //assume you need to check error messages

	if strings.Contains(strings.ToLower(title), "resolved") && todayDate(service, region, time) {
		runErrorMsg = false //no need to check error Messages, already resolved or didn't happen today
	}

	if runErrorMsg {
		for _, s := range errorMessages {
			if strings.Contains(title, s) {
				status = true
			}
		}
	}

	log.Printf("[DEBUG][STATUS CHECK] %s-%s: Returning Status Check: %v", service, region, status)

	return status
}

//todayDate - returns true if d is today's today
func todayDate(service, region, d string) bool {

	log.Printf("[DEBUG][TODAY DATE] %s-%s: Time: %s", service, region, d)

	//pubdate cames in as <pubDate>Thu, 22 Jun 2017 23:59:43 PDT</pubDate>

	now := time.Now()

	//sep := strings.Index(d, ",")
	//d2 := d[sep+1:]
	const awsForm = "Thu,  2 Jun 2017 23:59:43 PDT"
	compare, err := time.Parse(awsForm, d)
	if err != nil {
		log.Printf("[ERROR][TODAYS DATE] %s-%s: Parsing time Item: %s", service, region, err)
		return false
	}

	//todays day and year, this is an error message we should alert on
	if (now.Day() == compare.Day()) && (now.Year() == compare.Year()) {
		return true
	}

	return false
}
