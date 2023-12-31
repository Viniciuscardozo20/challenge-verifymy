package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"challenge-verifymy/customerror"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/viant/toolbox/format"
)

type UserReq struct {
	ID       string `json:"-"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type UserRes struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Address  string `json:"address" bson:"address"`
}

// Validate check if all fields are valid.
func (u *UserReq) Validate() error {
	return validator.New().Struct(u)
}

// Decode unmarshal User request data.
func (u *UserReq) Decode(body io.ReadCloser) error {
	if body == http.NoBody || body == nil {
		return customerror.ErrMissingBody
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Error(err, "failed to close reader body")
		}
	}(body)

	if err := json.NewDecoder(body).Decode(u); err != nil {
		return fmt.Errorf("failed to decode a user request: %w", err)
	}

	return nil
}

func (*UserReq) CheckFieldErrors(err error) map[string]string {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {

		messages := make(map[string]string)

		for _, fieldError := range fieldErrors {
			if _, ok = messages[fieldError.Field()]; !ok {
				field := format.CaseUpperCamel.Format(fieldError.Field(), format.CaseLowerUnderscore)

				messages[field] = fmt.Sprintf("%s is missing or invalid", field)
			}
		}

		return messages
	}

	return nil
}
