package models

import (
	"bytes"
	"challenge-verifymy/customerr"
	"encoding/json"
	"io"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateRequest(t *testing.T) {
	t.Run("Validate with success", func(t *testing.T) {
		userReq := &UserReq{
			Name:         "DummyName",
			Age:          10,
			Email:        "dummy@email.com",
			PasswordHash: "hashed_password",
			Address:      "dummy-address",
		}

		err := userReq.Validate()
		assert.NoError(t, err)
	})

	t.Run("Validate with missing field", func(t *testing.T) {
		userReq := &UserReq{
			Name:         "",
			Age:          10,
			Email:        "dummy@email.com",
			PasswordHash: "dummy-password",
			Address:      "dummy-address",
		}

		err := userReq.Validate()
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				assert.Equal(t, err.Field(), "Name")
			}
		}
	})

	t.Run("Validate with wrong field", func(t *testing.T) {
		userReq := &UserReq{
			Name:         "DummyName",
			Age:          10,
			Email:        "dummemail.com",
			PasswordHash: "dummy-password",
			Address:      "dummy-address",
		}

		err := userReq.Validate()
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				assert.Equal(t, err.Field(), "Email")
			}
		}
	})
}

func TestDecodeUserRequest(t *testing.T) {
	userReq := &UserReq{}

	t.Run("decode with empty body", func(t *testing.T) {
		err := userReq.Decode(nil)
		assert.Error(t, err, customerr.ErrMissingBody)
	})

	t.Run("decode with success", func(t *testing.T) {
		userReq := &UserReq{
			Name:         "DummyName",
			Age:          10,
			Email:        "dummy@email.com",
			PasswordHash: "hashed_password",
			Address:      "dummy-address",
		}

		data, err := json.Marshal(userReq)
		assert.NoError(t, err)

		reader := io.NopCloser(bytes.NewReader(data))
		err = userReq.Decode(reader)
		assert.NoError(t, err)
	})

	t.Run("decode with error", func(t *testing.T) {
		req := "data"

		data, err := json.Marshal(req)
		assert.NoError(t, err)

		reader := io.NopCloser(bytes.NewReader(data))

		err = userReq.Decode(reader)
		assert.ErrorContains(t, err, "failed to decode a user request")
	})
}
