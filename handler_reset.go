package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDB(context.Background())
	if err != nil {
		return fmt.Errorf("could't reset the database: %w", err)
	}

	fmt.Println("Reset was successful")
	return nil
}
