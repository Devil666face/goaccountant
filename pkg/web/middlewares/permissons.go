package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/gofiber/fiber/v2"
)

var ErrNotPermissions = fiber.ErrNotFound

func Admin(unit *web.Unit) error {
	if user, ok := unit.ViewCtx().Locals(web.UserKey).(models.User); ok {
		if user.Admin {
			return unit.ViewCtx().Next()
		}
	}
	return ErrNotPermissions
}
