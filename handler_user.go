package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rangaroo/gator-go/internal/database"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could't get users", err)
	}

	if len(users) == 0 {
		fmt.Println("No users")
		return nil
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user.Name)
		} else {
			fmt.Printf(" * %v\n", user.Name)
		}
	}

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user already exists\n")
	}


	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("could't create user: %w", err)
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could't set current user: %w", err)
	}

	fmt.Println("User created")
	printUser(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("could't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("could't set current user: %w", err)
	}

	fmt.Println("Username has been updated to:", s.cfg.CurrentUserName)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
