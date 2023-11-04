package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/gofiber/fiber/v2"
)

func Auth(uof *web.Uof) error {
	var (
		u   = models.User{}
		uID any
		err error
		ok  bool
	)
	if auth, err := uof.GetFromSession(web.AuthKey); auth == nil || err != nil {
		return uof.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if uID, err = uof.GetFromSession(web.UserID); uID == nil || err != nil {
		return uof.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if u.ID, ok = uID.(uint); !ok {
		return uof.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	if err := u.GetUser(uof.Database()); err != nil {
		return uof.ViewCtx().Status(fiber.StatusUnauthorized).
			RedirectToRoute("login", nil)
	}
	uof.ViewCtx().Locals(web.UserKey, u)
	return uof.ViewCtx().Next()
}
