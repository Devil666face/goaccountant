package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
)

const (
	AuthKey string = "authenticated"
	UserID  string = "user_id"
)

func AuthMiddleware(c *fiber.Ctx, _ *config.Config, _ *database.Database, s *session.Store) error {
	session, err := s.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if session.Get(AuthKey) == nil {
		return c.Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	uID := session.Get(UserID)
	if uID == nil {
		return c.Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	//Get user
	return c.Next()
}
