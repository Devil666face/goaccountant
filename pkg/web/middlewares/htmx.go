package middlewares

import "github.com/gofiber/fiber/v2"

const (
	Htmx      = "htmx"
	hxRequest = "Hx-Request"
)

func HtmxMiddleware(c *fiber.Ctx) error {
	c.Locals(Htmx, false)
	if _, ok := c.GetReqHeaders()[hxRequest]; ok {
		c.Locals(Htmx, true)
	}
	return c.Next()
}
