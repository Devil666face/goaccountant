package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx, _ *config.Config, _ *database.Database, _ *session.Store) error {
	return ViewCtx{c}.RenderWithCtx("login", fiber.Map{}, "base")
}
