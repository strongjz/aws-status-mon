package main

import (
	"github.com/strongjz/aws-status-mon/rss"
	"log"
	//"math/rand"
	//"time"
)

//var goroutineDelta = make(chan int)

func main() {

	//Get the Feed Data
	rssFeed, err := rss.GetFeed()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of Feeds %d\n", len(rssFeed))

	/*

		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		inum := random.Intn(len(rssFeed))

		log.Printf("Random Feed %s\n", rss.PrintFeed(rssFeed[inum]))

			newFeed := rss.NewFeed()

			newFeed.Service = "sqs"
			newFeed.Region = "us-east-1"
			newFeed.PollInt = 60
			newFeed.URL = "https://status.aws.amazon.com/rss/sqs-us-east-1.rss"
			//test polling one feed
	*/

	rss.PollFeed(rssFeed)

}
