package main

import (
	"log"
	"os"

	"github.com/rangaroo/gator-go/internal/config"
)

type state struct {
	cfg    *config.Config
}

func main() {
	cfg, err := config.Read()	
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmd := command {
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
