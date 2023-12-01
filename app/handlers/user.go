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
	ctx     context.Context
}

// SetUserRoutes creates user routes
func (h *handler) SetUserRoutes(ctx context.Context, app *fiber.App, service ports.UserService) {
	app.Get("/api/v1/:id", h.read)
	app.Get("/api/v1/", h.readAll)
	app.Post("/api/v1/", h.create)
	app.Put("/api/v1/:id", h.update)
	app.Delete("/api/v1/:id", h.delete)
}

func (h *handler) read(c *fiber.Ctx) error {
	id := c.Params("id")

	userRes, err := h.service.Read(h.ctx, id)
	if err != nil {
		if errors.Is(err, customerr.ErrFailedToFindDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "cannot find a user with that Id"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userRes}})
}

func (h *handler) readAll(c *fiber.Ctx) error {
	users, err := h.service.ReadAll(h.ctx)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "users": users})
}

func (h *handler) create(c *fiber.Ctx) error {
	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := h.service.Create(h.ctx, userReq); err != nil {
		if errors.Is(err, customerr.ErrFailedToInsertDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "failed to create a user"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "created"})
}

func (h *handler) update(c *fiber.Ctx) error {
	id := c.Params("id")

	var userReq *models.UserReq

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := userReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userRes, err := h.service.Update(h.ctx, id, userReq)
	if err != nil {
		if errors.Is(err, customerr.ErrFailedToUpdateDocument) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "failed to update a user"})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": userRes}})
}

func (h *handler) delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.Delete(h.ctx, id)
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
		ctx:     ctx,
		service: service,
	}
}
