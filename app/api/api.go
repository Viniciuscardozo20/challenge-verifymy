package api

import (
	"context"

	"challenge-verifymy/app/handlers"
	"challenge-verifymy/config"
	"challenge-verifymy/core/ports"
	"challenge-verifymy/core/services"
	"challenge-verifymy/infrastructure/mongodb"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Api struct {
	config   config.Config
	service  ports.UserService
	Shutdown func()
}

const userRepository = "user"

// New creates a new API
func New(ctx context.Context, cfg config.Config) (ap *Api, err error) {
	ap = &Api{config: cfg}

	db, err := mongodb.New(ctx, cfg.DBURI, cfg.DBName)
	if err != nil {
		return nil, err
	}

	userRepo := db.GetRepository(userRepository)
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

		userHandler.SetUserRoutes(app)

		log.Printf("Listening on addres and port %s", a.config.Api)

		go a.shutdown(ctx, app)

		return app.Listen(a.config.Api)
	}
}

func (a *Api) shutdown(ctx context.Context, app *fiber.App) {
	<-ctx.Done()
	log.Printf("Shutting down API gracefully...")
	app.Shutdown()
	a.Shutdown()
}
