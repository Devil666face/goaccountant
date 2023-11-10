package middlewares

import (
	"strings"

	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func AllowHost(unit *web.Unit) error {
	if host, ok := unit.ViewCtx().GetReqHeaders()[web.Host]; ok {
		if strings.Contains(host[0], unit.Config().AllowHost) {
			return unit.ViewCtx().Next()
		}
	}
	return fiber.ErrBadRequest
}
