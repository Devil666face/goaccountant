package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ViewCtx struct {
	*fiber.Ctx
}

func (c ViewCtx) Csrf() string {
	return c.Locals(middlewares.Csrf).(string)
}

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"c": ViewCtx{c}})
}
