package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2"
)

func Index(uof *web.Uof) error {
	return uof.ViewCtx().RenderWithCtx("login", fiber.Map{
		"Title": "Index",
	}, "index")
}
