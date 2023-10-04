package web

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/web/routes"
	"github.com/gofiber/fiber/v2"
)

type WebApp struct {
	app         *fiber.App
	logger      func(*fiber.Ctx) error
	static      func(*fiber.Ctx) error
	media       *Media
	middlewares []func(*fiber.Ctx) error
	config      *config.Config
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
		config:      config.New(),
	}
}

func (wa *WebApp) Listen() error {
	if wa.config.UseTls {
		return wa.listenTLS()
	}
	return wa.listenNoTLS()
}

func (wa *WebApp) listenTLS() error {
	go wa.redirectServer()
	if err := wa.app.ListenTLS(wa.config.ConnectHttps, wa.config.TlsCrt, wa.config.TlsKey); err != nil {
		return err
	}
	return nil
}

func (wa *WebApp) listenNoTLS() error {
	if err := wa.app.Listen(wa.config.ConnectHttp); err != nil {
		return err
	}
	return nil
}

func (wa *WebApp) redirectServer() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect(wa.config.HttpsRedirect)
	})
	if err := app.Listen(wa.config.ConnectHttp); err != nil {
		panic(err)
	}
}
