package middlewares

import (
	"strings"

	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func AllowedHostMiddleware(uof *web.Uof) error {
	if host, ok := uof.Ctx().GetReqHeaders()[web.Host]; ok {
		if strings.Contains(host, uof.Config().AllowHost) {
			return uof.Ctx().Next()
		}
	}
	return fiber.ErrBadRequest
}
