package main

import (
	"github.com/strongjz/aws-status-mon/rss"
	"log"
)

func main() {

	rssFeed, err := rss.GetFeed()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of Feeds %d\n", len(rssFeed))

	rss.PollFeed(rssFeed)

}
