package main

import (
	"fmt"
	"time"
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rangaroo/gator-go/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if (err == nil) {
		fmt.Println("user with that already exists")
		os.Exit(1)
	}

	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s was created\n", user.Name)
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Created at: %s\n", user.CreatedAt)
	fmt.Printf("Updated at: %s\n", user.UpdatedAt)

	return nil
}
