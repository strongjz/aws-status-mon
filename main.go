package main

import (
	"github.com/strongjz/aws-status-mon/rss"
	"log"
)

func main() {

	r := rss.NewRss()

	err := r.GetFeed()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[INF0] Number of Feeds %d\n", len(r.Feed))

	r.PollFeed()
}
