package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLogger() func(*fiber.Ctx) error {
	return logger.New()
}
