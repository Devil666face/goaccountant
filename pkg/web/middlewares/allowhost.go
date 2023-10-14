package middlewares

import (
	"strings"

	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
)

const (
	Host = "Host"
)

func AllowedHostMiddleware(c *fiber.Ctx, cfg *config.Config, _ *database.Database, _ *session.Store) error {
	if host, ok := c.GetReqHeaders()[Host]; ok {
		if strings.Contains(host, cfg.AllowHost) {
			return c.Next()
		}
	}
	return fiber.ErrBadRequest
}
