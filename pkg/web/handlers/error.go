package handlers

import (
	"fmt"

	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	//nolint:errorlint //Because crash page for any errors, if not convertation to fiber.Error - return 500
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return web.NewViewCtx(c).RenderWithCtx("error", fiber.Map{
		"Statuscode": code,
		"Error":      err.Error(),
		"Title":      fmt.Sprintf("Error %d", code),
	}, "base")
}
