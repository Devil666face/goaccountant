package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2"
)

func Htmx(uof *web.Uof) error {
	uof.ViewCtx().Locals(web.Htmx, false)
	if _, ok := uof.ViewCtx().GetReqHeaders()[web.HxRequest]; ok {
		uof.ViewCtx().Locals(web.Htmx, true)
	}
	return uof.ViewCtx().Next()
}

func HxOnly(uof *web.Uof) error {
	if uof.ViewCtx().IsHtmx() {
		return uof.ViewCtx().Next()
	}
	return fiber.ErrBadRequest
}
