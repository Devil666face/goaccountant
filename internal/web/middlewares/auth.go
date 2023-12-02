package middlewares

import (
	"github.com/Devil666face/goaccountant/internal/models"
	"github.com/Devil666face/goaccountant/internal/web/handlers"

	"github.com/gofiber/fiber/v2"
)

func Auth(h *handlers.Handler) error {
	var (
		u   = models.User{}
		uID any
		err error
		ok  bool
	)
	if auth, err := h.GetFromSession(handlers.AuthKey); auth == nil || err != nil {
		return h.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if uID, err = h.GetFromSession(handlers.UserID); uID == nil || err != nil {
		return h.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if u.ID, ok = uID.(uint); !ok {
		return h.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if err := u.Get(h.Database()); err != nil {
		return h.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	h.ViewCtx().Locals(handlers.UserKey, u)
	return h.ViewCtx().Next()
}

func AlreadyLogin(h *handlers.Handler) error {
	auth, err := h.GetFromSession(handlers.AuthKey)
	if auth, ok := auth.(bool); auth && ok && err == nil {
		return h.ViewCtx().RedirectToRoute("index", nil)
	}
	return h.ViewCtx().Next()
}
