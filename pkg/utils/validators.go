package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

var (
	policy                       = bluemonday.StrictPolicy()
	Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
)

var (
	ErrEmailRequired     = fiber.NewError(fiber.StatusBadRequest, "Email is required")
	ErrEmailIncorrect    = fiber.NewError(fiber.StatusBadRequest, "Incorrect email")
	ErrPasswordMissmatch = fiber.NewError(fiber.StatusBadRequest, "Passwords mismatch")
	ErrPasswordRequired  = fiber.NewError(fiber.StatusBadRequest, "Password is required")
	ErrPasswordShort     = fiber.NewError(fiber.StatusBadRequest, "Password is too short")
)

func validInput(in string) bool {
	return len([]rune(in)) == len([]rune(policy.Sanitize(in)))
}

func ValidateInputs(fields ...string) bool {
	for _, in := range fields {
		if !validInput(in) {
			return false
		}
	}
	return true
}

func ValidateInput(field string) bool {
	return validInput(field)
}

func SwitchUserValidateError(err error) error {
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Field() {
		case "Email":
			switch e.Tag() {
			case "required":
				return ErrEmailRequired
			case "email":
				return ErrEmailIncorrect
			}
		case "Password":
			switch e.Tag() {
			case "required":
				return ErrPasswordRequired
			case "min":
				return ErrPasswordShort
			}
		case "PasswordConfirm":
			switch e.Tag() {
			case "required":
				return ErrPasswordRequired
			case "eqfield":
				return ErrPasswordMissmatch
			}
		}
	}
	return err
}
