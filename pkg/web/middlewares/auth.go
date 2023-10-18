package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(uof *web.Uof) error {
	sess, err := uof.Store().Get(uof.FiberCtx())
	if err != nil {
		return uof.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if sess.Get(web.AuthKey) == nil {
		return uof.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	uID := sess.Get(web.UserID)
	if uID == nil {
		return uof.Ctx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	// Get user
	return uof.Ctx().Next()
}
