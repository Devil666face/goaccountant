package web

import (
	"github.com/Devil666face/goaccountant/pkg/web/routes"
	"github.com/gofiber/fiber/v2"
)

type WebApp struct {
	app         *fiber.App
	logger      func(*fiber.Ctx) error
	static      func(*fiber.Ctx) error
	media       *Media
	middlewares []func(*fiber.Ctx) error
}

func New() *WebApp {
	wa := Init()
	wa.app.Use(wa.logger)
	wa.app.Use(routes.StaticPrefix, wa.static)
	wa.app.Static(routes.MediaPrefix, wa.media.path, wa.media.handler)
	for _, m := range wa.middlewares {
		wa.app.Use(m)
	}
	return wa
}

func Init() *WebApp {
	return &WebApp{
		app: fiber.New(
			fiber.Config{
				ErrorHandler: nil,
				Views:        NewViews(),
			},
		),
		logger:      NewLogger(),
		static:      NewStatic(),
		media:       NewMedia(),
		middlewares: NewMiddlewares(),
	}
}
