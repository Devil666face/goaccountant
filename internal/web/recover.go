package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRecover() func(*fiber.Ctx) error {
	return recover.New()
}
