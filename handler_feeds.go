package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could't get users", err)
	}

	for i, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("could't get user: %w", err)
		}
		
		fmt.Printf("Feed %d\n", i)
		fmt.Println()
		fmt.Printf("Name:       %s\n", feed.Name)
		fmt.Printf("URL:        %s\n", feed.Url)
		fmt.Printf("Created by: %s\n", user.Name)
		
		fmt.Println("======================")
	}

	return nil
}
