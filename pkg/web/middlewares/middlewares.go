package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/web"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Logger(unit *web.Unit) error {
	return logger.New()(unit.Ctx())
}

func Recover(unit *web.Unit) error {
	return recover.New()(unit.Ctx())
}

func SecureHeaders(unit *web.Unit) error {
	// XSSProtection:             "0",
	// ContentTypeNosniff:        "nosniff",
	// XFrameOptions:             "SAMEORIGIN",
	// ReferrerPolicy:            "no-referrer",
	// CrossOriginEmbedderPolicy: "require-corp",
	// CrossOriginOpenerPolicy:   "same-origin",
	// CrossOriginResourcePolicy: "same-origin",
	// OriginAgentCluster:        "?1",
	// XDNSPrefetchControl:       "off",
	// XDownloadOptions:          "noopen",
	// XPermittedCrossDomain:     "none",
	return helmet.New(helmet.Config{
		XSSProtection:  "1",
		ReferrerPolicy: "same-origin",
	})(unit.Ctx())
}

func Compress(unit *web.Unit) error {
	return compress.New()(unit.Ctx())
}

func EncryptCookie(unit *web.Unit) error {
	return encryptcookie.New(encryptcookie.Config{
		Key: unit.Config().CookieKey,
	})(unit.Ctx())
}

func Limiter(unit *web.Unit) error {
	return limiter.New(limiter.Config{
		Storage: unit.Storage(),
		Max:     unit.Config().MaxQueryPerMinute,
	})(unit.Ctx())
}
