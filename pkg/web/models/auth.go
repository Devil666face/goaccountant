package models

import (
	"github.com/Devil666face/goaccountant/pkg/web/validators"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const WrongLoginData = "Mismatch username or password"

var (
	ErrUserNotFound  = fiber.NewError(fiber.StatusForbidden, WrongLoginData)
	ErrWrongPassword = fiber.NewError(fiber.StatusForbidden, WrongLoginData)
)

func (u *User) LoginValidate(db *gorm.DB, v *validators.Validator, password string) error {
	if !v.ValidateInputs(u.Email, u.Password, u.PasswordConfirm) {
		return fiber.ErrInternalServerError
	}
	if !u.IsFound(db) {
		return ErrUserNotFound
	}
	if !u.ComparePassword(password) {
		return ErrWrongPassword
	}
	return nil
}
