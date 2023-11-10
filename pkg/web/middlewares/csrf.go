package middlewares

import (
	"time"

	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func Csrf(unit *web.Unit) error {
	// Optional. Default: "header:X-CSRF-Token"
	return csrf.New(csrf.Config{
		Storage: unit.Storage(),
		// KeyLookup:  "form:csrf",
		CookieName: "csrf",
		// CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     web.Csrf,
		CookieHTTPOnly: true,
		SingleUseToken: true,
	})(unit.Ctx())
}
