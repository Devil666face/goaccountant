package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const WrongLoginData = "Mismatch username or password"

var (
	ErrUserNotFound  = fiber.NewError(fiber.StatusForbidden, WrongLoginData)
	ErrWrongPassword = fiber.NewError(fiber.StatusForbidden, WrongLoginData)
)

func (u *User) LoginValidate(db *gorm.DB, password string) error {
	if !u.validateInput() {
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