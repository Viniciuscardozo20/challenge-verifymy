package handlers

import (
	"context"
	"errors"

	"challenge-verifymy/common"
	"challenge-verifymy/core/models"
	"challenge-verifymy/core/ports"
	"challenge-verifymy/customerror"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type handler struct {
	service ports.UserService
}

// SetUserRoutes creates user routes
func (h *handler) SetUserRoutes(ctx context.Context, app *fiber.App) {
	app.Get("/api/v1/user/:id", func(c *fiber.Ctx) error {
		return h.read(c, ctx)
	})
	app.Get("/api/v1/user", func(c *fiber.Ctx) error {
		return h.readAll(c, ctx)
	})
	app.Post("/api/v1/user", func(c *fiber.Ctx) error {
		return h.create(c, ctx)
	})
	app.Put("/api/v1/user/:id", func(c *fiber.Ctx) error {
		return h.update(c, ctx)
	})
	app.Delete("/api/v1/user/:id", func(c *fiber.Ctx) error {
		return h.delete(c, ctx)
	})
}

func (h *handler) read(c *fiber.Ctx, ctx context.Context) error {
	id := c.Params("id")

	log.Info("Received request to read user with ID:", id)

	userRes, err := h.service.Read(ctx, id)
	if err != nil {
		if errors.Is(err, customerror.ErrFailedToFindDocument) {
			log.Warn("User not found with ID:", id)
			return c.Status(fiber.StatusNotFound).JSON(common.Data{Status: "fail", Message: "cannot find a user with that Id"})
		}

		log.Error("Failed to read user with ID:", id, " Error:", err)
		return c.Status(fiber.StatusBadGateway).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	log.Info("Successfully read user with ID:", id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userRes}})
}

func (h *handler) readAll(c *fiber.Ctx, ctx context.Context) error {
	log.Info("Received request to read all users")

	users, err := h.service.ReadAll(ctx)
	if err != nil {
		log.Error("Failed to read all users. Error:", err)
		return c.Status(fiber.StatusBadGateway).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	log.Info("Successfully read all users")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "users": users})
}

func (h *handler) create(c *fiber.Ctx, ctx context.Context) error {
	log.Info("Received request to create a user")

	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		log.Error("Failed to parse request body. Error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		log.Warn("Invalid user request. Error:", err)

		if fieldErrs := userReq.CheckFieldErrors(err); fieldErrs != nil {
			var errs []common.Data

			for _, message := range fieldErrs {
				errs = append(errs, common.Data{
					Status:  "fail",
					Message: message,
				})
			}

			return c.Status(fiber.StatusBadRequest).JSON(errs)
		}

		return c.Status(fiber.StatusBadRequest).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	user, err := h.service.Create(ctx, userReq)
	if err != nil {
		if errors.Is(err, customerror.ErrUserAlreadyExists) {
			log.Error("Failed to create a user. Error:", err)
			return c.Status(fiber.StatusConflict).JSON(common.Data{Status: "fail", Message: err.Error()})
		}

		log.Error("Failed to create a user. Error:", err)
		return c.Status(fiber.StatusBadGateway).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	log.Info("Successfully created a user")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func (h *handler) update(c *fiber.Ctx, ctx context.Context) error {
	log.Info("Received request to update a user")

	id := c.Params("id")

	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		log.Error("Failed to parse request body. Error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		log.Warn("Invalid user request. Error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userRes, err := h.service.Update(ctx, id, userReq)
	if err != nil {
		if errors.Is(err, customerror.ErrFailedToUpdateDocument) {
			log.Error("Failed to update a user. Error:", err)
			return c.Status(fiber.StatusNotFound).JSON(common.Data{Status: "fail", Message: "failed to update a user"})
		}

		log.Error("Failed to update a user. Error:", err)
		return c.Status(fiber.StatusBadGateway).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	log.Info("Successfully updated a user")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userRes}})
}

func (h *handler) delete(c *fiber.Ctx, ctx context.Context) error {
	log.Info("Received request to delete a user")

	id := c.Params("id")

	err := h.service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, customerror.ErrFailedToFindDocument) {
			log.Warn("Cannot find a user with the specified ID. Error:", err)
			return c.Status(fiber.StatusNotFound).JSON(common.Data{Status: "fail", Message: "cannot find a user with that Id"})
		}

		log.Error("Failed to delete a user. Error:", err)
		return c.Status(fiber.StatusBadGateway).JSON(common.Data{Status: "fail", Message: err.Error()})
	}

	log.Info("Successfully deleted a user")

	return c.SendStatus(fiber.StatusNoContent)
}

func NewUserHandler(ctx context.Context, service ports.UserService) *handler {
	return &handler{
		service: service,
	}
}
