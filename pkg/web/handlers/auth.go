package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web/view"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx, _ *config.Config, _ *database.Database, _ *session.Store) error {
	return view.New(c).RenderWithCtx("login", fiber.Map{
		"Title": "Login",
	}, "base")
}
