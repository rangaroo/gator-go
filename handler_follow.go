package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rangaroo/gator-go/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	userName := s.cfg.CurrentUserName
	url := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("could't get user: %w", err)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could't get feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could't create feed follow record: %w", err)
	}

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("could't get user: %w", err)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could't get following feeds: %w", err)
	}

	fmt.Printf("Feeds followed by %s:\n", userName)

	for _, item := range following {
		fmt.Printf(" - %s\n", item.FeedName)
	}

	return nil
}
