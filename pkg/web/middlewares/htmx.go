package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
)

func HtmxMiddleware(uof *web.Uof) error {
	uof.Ctx().Locals(web.Htmx, false)
	if _, ok := uof.Ctx().GetReqHeaders()[web.HxRequest]; ok {
		uof.Ctx().Locals(web.Htmx, true)
	}
	return uof.Ctx().Next()
}
