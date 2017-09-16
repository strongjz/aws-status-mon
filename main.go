package main

import (
	"github.com/strongjz/aws-status-mon/rss"
	"log"
	"math/rand"
	"time"
)

func main() {

	//Get the Feed Data
	rssFeed, err := rss.GetFeed()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of Feeds %d\n", len(rssFeed))

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	i := random.Intn(len(rssFeed))

	log.Printf("Random Feed Region - %s , Service - %s\n", rssFeed[i].Region, rssFeed[i].Service)

}
