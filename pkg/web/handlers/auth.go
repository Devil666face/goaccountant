package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func Login(uof *web.Uof) error {
	return uof.Ctx().RenderWithCtx("login", fiber.Map{
		"Title": "Login",
	}, "base")
}
