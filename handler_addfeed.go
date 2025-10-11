package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rangaroo/gator-go/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <url>", cmd.Name)
	}
	
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could't get the current user: %w", err)
	}

	name   := cmd.Args[0]
	url    := cmd.Args[1]
	userId := user.ID

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    userId,
	})
	if err != nil {
		return fmt.Errorf("could't create feed: %w", err)
	}

	fmt.Println("Feed created")
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * UserID:  %v\n", feed.UserID)
}
