package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2"
)

func Htmx(unit *web.Unit) error {
	unit.ViewCtx().Locals(web.Htmx, false)
	if _, ok := unit.ViewCtx().GetReqHeaders()[web.HxRequest]; ok {
		unit.ViewCtx().Locals(web.Htmx, true)
	}
	return unit.ViewCtx().Next()
}

func HxOnly(unit *web.Unit) error {
	if unit.ViewCtx().IsHtmx() {
		return unit.ViewCtx().Next()
	}
	return fiber.ErrBadRequest
}
