package middlewares

import (
	"fmt"

	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/gofiber/fiber/v2"
)

const (
	AuthKey string = "authenticated"
	UserID  string = "user_id"
)

func AuthMiddleware(s *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println("exec")
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
}
