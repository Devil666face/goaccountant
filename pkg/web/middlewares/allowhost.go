package middlewares

import (
	"strings"

	"github.com/Devil666face/goaccountant/pkg/config"

	"github.com/gofiber/fiber/v2"
)

const (
	Host = "Host"
)

func AllowedHostMiddleware(config *config.Config) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if host, ok := c.GetReqHeaders()[Host]; ok {
			if strings.Contains(host, config.AllowHost) {
				return c.Next()
			}
		}
		return fiber.ErrBadRequest
	}
}
