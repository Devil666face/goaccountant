package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
)

const (
	Htmx      = "htmx"
	hxRequest = "Hx-Request"
)

func HtmxMiddleware(c *fiber.Ctx, _ *config.Config, _ *database.Database, _ *session.Store) error {
	c.Locals(Htmx, false)
	if _, ok := c.GetReqHeaders()[hxRequest]; ok {
		c.Locals(Htmx, true)
	}
	return c.Next()
}
