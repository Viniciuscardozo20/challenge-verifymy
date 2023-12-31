package api

import (
	"context"

	"challenge-verifymy/app/handlers"
	"challenge-verifymy/config"
	"challenge-verifymy/core/ports"
	"challenge-verifymy/core/services"
	"challenge-verifymy/infrastructure/mongodb"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Api struct {
	config   config.Config
	service  ports.UserService
	Shutdown func()
}

const collName = "user"

// New creates a new API
func New(ctx context.Context, cfg config.Config) (ap *Api, err error) {
	ap = &Api{config: cfg}

	db, err := mongodb.New(ctx, cfg.DBURI, cfg.DBName)
	if err != nil {
		return nil, err
	}

	userRepo, err := db.GetRepository(ctx, collName)
	if err != nil {
		return nil, err
	}

	ap.service = services.NewUserService(userRepo)

	shutdown := func() {
		<-db.Disconnected()
	}

	ap.Shutdown = shutdown

	return ap, nil
}

// Run API
func (a *Api) Run(ctx context.Context, cancel context.CancelFunc) func() error {
	return func() error {
		defer cancel()

		app := fiber.New()

		userHandler := handlers.NewUserHandler(ctx, a.service)

		userHandler.SetUserRoutes(ctx, app)

		log.Info("Listening on address and port")

		go a.shutdown(ctx, app)

		return app.Listen(a.config.Api)
	}
}

func (a *Api) shutdown(ctx context.Context, app *fiber.App) {
	<-ctx.Done()
	log.Infof("Shutting down API gracefully...")
	if err := app.Shutdown(); err != nil {
		log.Infof("Error shutting down API: %v", err)
	}
	a.Shutdown()
}
