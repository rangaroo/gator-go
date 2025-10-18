package main

import (
	_ "github.com/lib/pq"
	"database/sql"
)

import (
	"context"
	"log"
	"os"
	"github.com/rangaroo/gator-go/internal/config"
	"github.com/rangaroo/gator-go/internal/database"
)

type state struct {
	cfg    *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()	
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}


	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)
	
	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	// Register courses
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
