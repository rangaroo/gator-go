package main

import (
	"fmt"
	"log"
	"context"
	"time"

	"github.com/rangaroo/gator-go/internal/database"

)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs -> Example: 1s | 1m | 1h>", cmd.Name)
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could't parse the time from input: %w", err)
	}

	fmt.Printf("Collecting feeds every %v...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("could't get next feed", err)
		return
	}
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Errorf("could't mark the feed as fetched: %w", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("could't fetch feed %s: %w", feed.Name, err)
		return
	}

	fmt.Printf("Printing the items in feed: %s\n", feedData.Channel.Title)
	for _, item := range feedData.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	fmt.Println("==============================================================")
}
