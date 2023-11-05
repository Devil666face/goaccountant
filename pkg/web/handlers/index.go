package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2"
)

func Index(uof *web.Uof) error {
	// uof.ViewCtx().SetClientRefresh()
	return uof.ViewCtx().RenderWithCtx("index", fiber.Map{
		"Title": "Index",
	}, "base")
}
