package handlers

import (
	"challenge-verifymy/core/models"
	"challenge-verifymy/core/ports"
	"challenge-verifymy/customerr"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ports.UserService
}

// SetUserRoutes creates user routes
func (h *handler) SetUserRoutes(ctx context.Context, app *fiber.App) {
	app.Get("/api/v1/:id", func(c *fiber.Ctx) error {
		return h.read(c, ctx)
	})
	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return h.readAll(c, ctx)
	})
	app.Post("/api/v1/", func(c *fiber.Ctx) error {
		return h.create(c, ctx)
	})
	app.Put("/api/v1/:id", func(c *fiber.Ctx) error {
		return h.update(c, ctx)
	})
	app.Delete("/api/v1/:id", func(c *fiber.Ctx) error {
		return h.delete(c, ctx)
	})
}

func (h *handler) read(c *fiber.Ctx, ctx context.Context) error {
	id := c.Params("id")

	userRes, err := h.service.Read(ctx, id)
	if err != nil {
		if errors.Is(err, customerr.ErrFailedToFindDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "cannot find a user with that Id"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userRes}})
}

func (h *handler) readAll(c *fiber.Ctx, ctx context.Context) error {
	users, err := h.service.ReadAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "users": users})
}

func (h *handler) create(c *fiber.Ctx, ctx context.Context) error {
	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := h.service.Create(ctx, userReq); err != nil {
		if errors.Is(err, customerr.ErrFailedToInsertDocument) {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "failed to create a user"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "created"})
}

func (h *handler) update(c *fiber.Ctx, ctx context.Context) error {
	id := c.Params("id")

	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userRes, err := h.service.Update(ctx, id, userReq)
	if err != nil {
		if errors.Is(err, customerr.ErrFailedToUpdateDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "failed to update a user"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": userRes}})
}

func (h *handler) delete(c *fiber.Ctx, ctx context.Context) error {
	id := c.Params("id")

	err := h.service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, customerr.ErrFailedToFindDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "cannot find a user with that Id"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func NewUserHandler(ctx context.Context, service ports.UserService) *handler {
	return &handler{
		service: service,
	}
}
