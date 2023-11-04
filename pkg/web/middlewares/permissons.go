package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/gofiber/fiber/v2"
)

var ErrNotPermissions = fiber.ErrNotFound

func Admin(uof *web.Uof) error {
	if user, ok := uof.ViewCtx().Locals(web.UserKey).(models.User); ok {
		if user.Admin {
			return uof.ViewCtx().Next()
		}
	}
	return ErrNotPermissions
}
