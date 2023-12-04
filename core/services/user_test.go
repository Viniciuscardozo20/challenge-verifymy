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

func TestRead(t *testing.T) {
	id := "123"

	t.Run("successful read", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		res := &models.UserRes{}

		userRepositoryMock.EXPECT().Read(ctx, id, res).Return(nil)

		service := NewUserService(userRepositoryMock)

		user, err := service.Read(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("failed read", func(t *testing.T) {
		ctx := context.Background()

		res := &models.UserRes{}

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().Read(ctx, id, res).Return(errors.New("some error"))

		service := NewUserService(userRepositoryMock)

		_, err := service.Read(ctx, id)
		assert.Error(t, err)
	})
}

func TestReadAll(t *testing.T) {
	t.Run("successful readall", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		res := &[]models.UserRes{}

		userRepositoryMock.EXPECT().ReadAll(ctx, res).Return(nil)

		service := NewUserService(userRepositoryMock)

		user, err := service.ReadAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("failed readall", func(t *testing.T) {
		ctx := context.Background()

		res := &[]models.UserRes{}

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().ReadAll(ctx, res).Return(errors.New("some error"))

		service := NewUserService(userRepositoryMock)

		_, err := service.ReadAll(ctx)
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	id := "123"

	req := &models.UserReq{
		Name:     "dummy-name",
		Age:      12,
		Email:    "dummy@email.com",
		Password: "1234",
		Address:  "dummy-add",
	}

	t.Run("successful update", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		res := &models.UserRes{}

		userRepositoryMock.EXPECT().Update(ctx, id, req, res).Return(nil)

		service := NewUserService(userRepositoryMock)

		user, err := service.Update(ctx, id, req)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("failed update", func(t *testing.T) {
		ctx := context.Background()

		res := &models.UserRes{}

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().Update(ctx, id, req, res).Return(errors.New("some error"))

		service := NewUserService(userRepositoryMock)

		_, err := service.Update(ctx, id, req)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	id := "123"

	t.Run("successful delete", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().Delete(ctx, id).Return(nil)

		service := NewUserService(userRepositoryMock)

		err := service.Delete(ctx, id)
		assert.NoError(t, err)
	})

	t.Run("failed delete", func(t *testing.T) {
		ctx := context.Background()

		userRepositoryMock := testutil.NewMockUserRepository(t)

		userRepositoryMock.EXPECT().Delete(ctx, id).Return(errors.New("some error"))

		service := NewUserService(userRepositoryMock)

		err := service.Delete(ctx, id)
		assert.Error(t, err)
	})
}
