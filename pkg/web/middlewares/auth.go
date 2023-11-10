package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

func Auth(unit *web.Unit) error {
	var (
		u   = models.User{}
		uID any
		err error
		ok  bool
	)
	if auth, err := unit.GetFromSession(web.AuthKey); auth == nil || err != nil {
		return unit.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if uID, err = unit.GetFromSession(web.UserID); uID == nil || err != nil {
		return unit.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if u.ID, ok = uID.(uint); !ok {
		return unit.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if err := u.Get(unit.Database()); err != nil {
		return unit.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	unit.ViewCtx().Locals(web.UserKey, u)
	return unit.ViewCtx().Next()
}

func AlreadyLogin(unit *web.Unit) error {
	auth, err := unit.GetFromSession(web.AuthKey)
	if auth, ok := auth.(bool); auth && ok && err == nil {
		return unit.ViewCtx().RedirectToRoute("index", nil)
	}
	return unit.ViewCtx().Next()
}
