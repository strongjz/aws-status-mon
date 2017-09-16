package main

import (
	"github.com/strongjz/aws-status-mon/feed"
	"log"
	"math/rand"
	"time"
)

func main() {
	rssFeed, err := feed.GetFeed()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of Feeds %d\n", len(rssFeed))

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	i := random.Intn(len(rssFeed))

	log.Printf("Random Feed %s\n", rssFeed[i])

}
