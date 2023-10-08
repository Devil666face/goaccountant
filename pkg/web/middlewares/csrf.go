package middlewares

import (
	"time"

	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

const (
	Csrf = "csrf"
)

func CsrfMiddleware(s *session.SessionStore) func(c *fiber.Ctx) error {
	return csrf.New(csrf.Config{
		Storage:        s.Storage,
		KeyLookup:      "form:csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     Csrf,
	})
}
