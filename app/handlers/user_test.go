package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"challenge-verifymy/core/models"
	testutil "challenge-verifymy/core/ports/testutil"
	"challenge-verifymy/customerror"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadUser(t *testing.T) {
	app := fiber.New()

	mockUserService := testutil.NewMockUserService(t)

	handler := NewUserHandler(context.Background(), mockUserService)

	// Set up test route
	app.Get("/api/v1/:id", func(c *fiber.Ctx) error {
		return handler.read(c, context.Background())
	})

	tests := []struct {
		Name         string
		UserID       string
		ExpectedCode int
		MockFunc     func()
	}{
		{
			Name:         "Success",
			UserID:       "123",
			ExpectedCode: fiber.StatusOK,
			MockFunc: func() {
				expectedUser := &models.UserRes{ID: "123", Name: "vini"}
				mockUserService.EXPECT().Read(mock.Anything, "123").Return(expectedUser, nil)
			},
		},
		{
			Name:         "User Not Found",
			UserID:       "456",
			ExpectedCode: fiber.StatusNotFound,
			MockFunc: func() {
				mockUserService.EXPECT().Read(mock.Anything, "456").Return(nil, customerror.ErrFailedToFindDocument)
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.MockFunc()

			req := httptest.NewRequest("GET", "/api/v1/"+tt.UserID, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedCode, resp.StatusCode)
		})
	}
}

func TestCreateHandler(t *testing.T) {
	app := fiber.New()

	mockUserService := testutil.NewMockUserService(t)

	handler := NewUserHandler(context.Background(), mockUserService)

	app.Post("/api/v1/", func(c *fiber.Ctx) error {
		return handler.create(c, context.Background())
	})

	tests := []struct {
		Name         string
		RequestJSON  string
		ExpectedCode int
		MockFunc     func()
	}{
		{
			Name:         "Success",
			RequestJSON:  `{"name": "Vini", "age": 25, "email": "vini@email.com", "password": "pass123", "address": "123 Street"}`,
			ExpectedCode: fiber.StatusCreated,
			MockFunc: func() {
				expectedUser := &models.UserRes{ID: "123", Name: "Vini"}
				mockUserService.EXPECT().Create(mock.Anything, mock.Anything).Return(expectedUser, nil).Times(1)
			},
		},
		{
			Name:         "Invalid Request Body",
			RequestJSON:  `invalid_json`,
			ExpectedCode: fiber.StatusBadRequest,
		},
		{
			Name:         "Validation Error",
			RequestJSON:  `{"name": "", "age": 25, "email": "vini@email.com", "password": "pass123", "address": "123 Street"}`,
			ExpectedCode: fiber.StatusBadRequest,
		},
		{
			Name:         "User Already Exists",
			RequestJSON:  `{"name": "Vini", "age": 25, "email": "vini@email.com", "password": "pass123", "address": "123 Street"}`,
			ExpectedCode: fiber.StatusConflict,
			MockFunc: func() {
				mockUserService.EXPECT().Create(mock.Anything, mock.Anything).Return(nil, customerror.ErrUserAlreadyExists).Times(1)
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.MockFunc != nil {

				tt.MockFunc()
			}

			req := httptest.NewRequest("POST", "/api/v1/", bytes.NewBufferString(tt.RequestJSON))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedCode, resp.StatusCode)
		})
	}
}

func TestUpdateHandler(t *testing.T) {
	app := fiber.New()

	mockUserService := testutil.NewMockUserService(t)

	handler := NewUserHandler(context.Background(), mockUserService)

	app.Put("/api/v1/:id", func(c *fiber.Ctx) error {
		return handler.update(c, context.Background())
	})

	tests := []struct {
		Name         string
		RequestID    string
		RequestJSON  string
		ExpectedCode int
		MockFunc     func()
	}{
		{
			Name:         "Success",
			RequestID:    "123",
			RequestJSON:  `{"name": "Updated User", "age": 30, "email": "updated@email.com", "password": "updatedpass", "address": "456 Street"}`,
			ExpectedCode: fiber.StatusCreated,
			MockFunc: func() {
				expectedUser := &models.UserRes{ID: "123", Name: "Updated User"}
				mockUserService.EXPECT().Update(mock.Anything, "123", mock.Anything).Return(expectedUser, nil).Times(1)
			},
		},
		{
			Name:         "Invalid Request Body",
			RequestID:    "123",
			RequestJSON:  `invalid_json`,
			ExpectedCode: fiber.StatusBadRequest,
		},
		{
			Name:         "Validation Error",
			RequestID:    "123",
			RequestJSON:  `{"name": "", "age": 30, "email": "updated@email.com", "password": "updatedpass", "address": "456 Street"}`,
			ExpectedCode: fiber.StatusBadRequest,
		},
		{
			Name:         "User Not Found",
			RequestID:    "456",
			RequestJSON:  `{"name": "Updated User", "age": 30, "email": "updated@email.com", "password": "updatedpass", "address": "456 Street"}`,
			ExpectedCode: fiber.StatusNotFound,
			MockFunc: func() {
				mockUserService.EXPECT().Update(mock.Anything, "456", mock.Anything).Return(nil, customerror.ErrFailedToUpdateDocument).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.MockFunc != nil {
				tt.MockFunc()
			}

			req := httptest.NewRequest("PUT", "/api/v1/"+tt.RequestID, bytes.NewBufferString(tt.RequestJSON))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedCode, resp.StatusCode)
		})
	}
}

func TestReadAllHandler(t *testing.T) {
	app := fiber.New()

	mockUserService := testutil.NewMockUserService(t)

	handler := NewUserHandler(context.Background(), mockUserService)

	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return handler.readAll(c, context.Background())
	})

	tests := []struct {
		Name         string
		ExpectedCode int
		MockFunc     func()
	}{
		{
			Name:         "Success",
			ExpectedCode: fiber.StatusOK,
			MockFunc: func() {
				expectedUsers := []models.UserRes{{ID: "1", Name: "User 1"}, {ID: "2", Name: "User 2"}}
				mockUserService.EXPECT().ReadAll(mock.Anything).Return(&expectedUsers, nil).Times(1)
			},
		},
		{
			Name:         "Failure",
			ExpectedCode: fiber.StatusBadGateway,
			MockFunc: func() {
				mockUserService.EXPECT().ReadAll(mock.Anything).Return(nil, errors.New("some error")).Times(1)
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.MockFunc != nil {
				tt.MockFunc()
			}

			req := httptest.NewRequest("GET", "/api/v1/", nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedCode, resp.StatusCode)
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	app := fiber.New()

	mockUserService := testutil.NewMockUserService(t)

	handler := NewUserHandler(context.Background(), mockUserService)

	app.Delete("/api/v1/:id", func(c *fiber.Ctx) error {
		return handler.delete(c, context.Background())
	})

	tests := []struct {
		Name         string
		ExpectedCode int
		MockFunc     func()
	}{
		{
			Name:         "Success",
			ExpectedCode: fiber.StatusNoContent,
			MockFunc: func() {
				mockUserService.EXPECT().Delete(mock.Anything, mock.Anything).Return(nil).Times(1)
			},
		},
		{
			Name:         "UserNotFound",
			ExpectedCode: fiber.StatusNotFound,
			MockFunc: func() {
				mockUserService.EXPECT().Delete(mock.Anything, mock.Anything).Return(customerror.ErrFailedToFindDocument).Times(1)
			},
		},
		{
			Name:         "Failure",
			ExpectedCode: fiber.StatusBadGateway,
			MockFunc: func() {
				mockUserService.EXPECT().Delete(mock.Anything, mock.Anything).Return(errors.New("some error")).Times(1)
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.MockFunc != nil {
				tt.MockFunc()
			}

			req := httptest.NewRequest("DELETE", "/api/v1/123", nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedCode, resp.StatusCode)
		})
	}
}
