package main

import (
	"challenge-verifymy/app/api"
	"challenge-verifymy/config"
	"context"

	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("Failed to load environment variables! \n", err.Error())
	}

	var g multierror.Group
	ctx, cancel := context.WithCancel(context.Background())

	a, err := api.New(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	g.Go(a.Run(ctx, cancel))

	if err := g.Wait().ErrorOrNil(); err != nil {
		log.Fatal(err)
	}
}
