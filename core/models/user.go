package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/viant/toolbox/format"
)

type User struct {
	Name         string `json:"name" validate:"required"`
	Age          int    `json:"age" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"password" validate:"required"`
	Address      string `json:"address" validate:"required"`
}

// Validate check if all fields are valid.
func (u *User) Validate() error {
	return validator.New().Struct(u)
}

// CheckFieldErrors checks if the given error is a validation error and if so, it returns a map containing the wrong fields
func (*User) CheckFieldErrors(err error) map[string]string {
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
