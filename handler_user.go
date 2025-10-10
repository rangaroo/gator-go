package main

import (
	"fmt"
	"context"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("user with such name does not exist")
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("Username has been updated to:", s.cfg.CurrentUserName)
	return nil
}
