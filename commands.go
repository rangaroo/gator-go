package main

import (
	"errors"
)

type command struct {
	Name     string
	Args     []string
}

type commands struct {
	registry  map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.registry[cmd.Name]
	if !exists {
		return errors.New("Invalid command")
	}

	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}
