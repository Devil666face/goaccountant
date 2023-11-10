package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func Csrf(unit *web.Unit) error {
	// KeyLookup:         "header:" + "X-Csrf-Token",
	// CookieName:        "csrf_",
	// CookieSameSite:    "Lax",
	// Expiration:        1 * time.Hour,
	// KeyGenerator:      utils.UUIDv4,
	// ErrorHandler:      defaultErrorHandler,
	// Extractor:         CsrfFromHeader("X-Csrf-Token"),
	// SessionKey:        "fiber.csrf.token",
	// HandlerContextKey: "fiber.csrf.handler",
	return csrf.New(csrf.Config{
		Storage: unit.Storage(),
		// KeyLookup:  "form:csrf",
		ContextKey:     web.Csrf,
		CookieHTTPOnly: true,
		SingleUseToken: true,
	})(unit.Ctx())
}
