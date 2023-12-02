package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(h *Handler) error {
	// unit.ViewCtx().SetClientRefresh()
	return h.ViewCtx().RenderWithCtx("index", fiber.Map{
		"Title": "Index",
	}, "base")
}
