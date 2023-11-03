package middlewares

import (
	"strings"

	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func AllowHost(uof *web.Uof) error {
	if host, ok := uof.ViewCtx().GetReqHeaders()[web.Host]; ok {
		if strings.Contains(host[0], uof.Config().AllowHost) {
			return uof.ViewCtx().Next()
		}
	}
	return fiber.ErrBadRequest
}
