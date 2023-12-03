package main

import (
	"challenge-verifymy/app/api"
	"challenge-verifymy/common"
	"challenge-verifymy/config"
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/hashicorp/go-multierror"
)

func main() {
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("Failed to load environment variables! \n", err.Error())
	}

	log.SetLevel(common.ParseLogLevel(config.LogLevel))

	log.Info("Starting the application")

	var g multierror.Group
	ctx, cancel := context.WithCancel(context.Background())

	a, err := api.New(ctx, config)
	if err != nil {
		log.Fatal("Failed to create API instance. Error:", err)
	}

	g.Go(a.Run(ctx, cancel))

	if err := g.Wait().ErrorOrNil(); err != nil {
		log.Fatal("An error occurred while running the application. Error:", err)
	}

	log.Info("Application stopped")
}
