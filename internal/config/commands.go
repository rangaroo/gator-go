package config

import (
	"errors"
	"fmt"
)

type state struct {
	cfg    *Config
}

type command struct {
	name     string
	args     []string
}

func handlerLogin(s *state, cmd command) error {
	if len(command.args) != 1 {
		return errors.New("You must provide a username")
	}

	s.cfg.CurrentUserName = cmd.args[0]

	fmt.Println("Username has been updated to: ", s.cfg.CurrentUserName)
	return nil
}

type commands struct {
	cmd       map[string]func(*state, command) error
}

func (c *commands) run (s *state, cmd command) error {
	commandHandler, exists := c.cmd[cmd.name]
	if !exists {
		return errors.New("Invalid command")
	}

	return commandHandler(*state, cmd)
}
