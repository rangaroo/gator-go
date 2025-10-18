package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rangaroo/gator-go/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <url>", cmd.Name)
	}
	
	name   := cmd.Args[0]
	url    := cmd.Args[1]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could't create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could't create feed follow record: %w", err)
	}

	fmt.Println("Feed created")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("Feed follow record created")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=============================================")
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could't get feeds", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("could't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=============================================")
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Created: %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated: %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:     %v\n", feed.Url)
	fmt.Printf(" * User:    %v\n", user.Name)
}
