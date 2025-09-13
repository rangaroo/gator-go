package main

import (
	"fmt"

	"github.com/rangaroo/gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("rangaroo")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg.DBURL)
	fmt.Println(cfg.CurrentUserName)
}
