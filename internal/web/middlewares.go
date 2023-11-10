package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Use `openssl rand -base64 32` for get hash
const (
	cookieKey = "VtsTmAz5I7LUM3N2NA4J7eX1XC/gNzA8DUK1Ocssowo="
)

func NewLogger() func(*fiber.Ctx) error {
	return logger.New()
}

func NewRecover() func(*fiber.Ctx) error {
	return recover.New()
}

func NewHelmet() func(*fiber.Ctx) error {
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
	})
}

func NewCompress() func(*fiber.Ctx) error {
	return compress.New()
}

func NewEncryptCookie() func(*fiber.Ctx) error {
	return encryptcookie.New(encryptcookie.Config{
		Key: cookieKey,
	})
}
