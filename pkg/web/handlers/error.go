package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return ViewCtx{c}.RenderWithCtx("error", fiber.Map{
		"Statuscode": code,
		"Error":      err.Error(),
		"Title":      fmt.Sprintf("Error %d", code),
	}, "base")
}
