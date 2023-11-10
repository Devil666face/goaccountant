package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter(unit *web.Unit) error {
	return limiter.New(limiter.Config{
		Storage: unit.Storage(),
		Max:     unit.Config().MaxQueryPerMinute,
	})(unit.Ctx())
}
