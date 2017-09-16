package main

import (
	"fmt"
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

	fmt.Printf("Random Feed %s\n", rss.PrintFeed(rssFeed[i]))

}
