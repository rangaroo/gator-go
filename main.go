package main

import (
	"fmt"
	"log"
	"github.com/rangaroo/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("rangaroo")
	if err != nil {
		log.Fatalf("error while updating the current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Updated the config: %v\n", cfg)
}
