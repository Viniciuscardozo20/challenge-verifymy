package services

import (
	"challenge-verifymy/core/models"
	testutil "challenge-verifymy/core/ports/testutil"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	req := &models.UserReq{
		ID:       "1234",
		Name:     "dummy-name",
		Age:      12,
		Email:    "dummy@email.com",
		Password: "1234",
		Address:  "dummy-add",
	}

	t.Run("successful create", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		res := &models.UserRes{}

		userRepositoryMock.EXPECT().Save(ctx, req, res).Return(nil)

		service := NewUserService(userRepositoryMock)

		user, err := service.Create(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("failed create", func(t *testing.T) {
		ctx := context.Background()

		res := &models.UserRes{}

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().Save(ctx, req, res).Return(errors.New("failed"))

		service := NewUserService(userRepositoryMock)

		_, err := service.Create(ctx, req)
		assert.Error(t, err)
	})
}
