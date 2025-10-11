package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could't fetch feed: %w", err)
	}

    fmt.Printf("%+v\n", *feed)
	return nil
}
