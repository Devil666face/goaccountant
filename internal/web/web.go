package web

import (
	"log"

	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"
	"github.com/Devil666face/goaccountant/pkg/web/routes"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	app         *fiber.App
	logger      func(*fiber.Ctx) error
	static      func(*fiber.Ctx) error
	media       *Media
	config      *config.Config
	database    *database.Database
	middlewares *middlewares.Middlewares
	router      *routes.AppRouter
	session     *session.Store
}

func New() *App {
	a := Init()

	a.database = database.New(a.config, []interface{}{})
	a.session = session.New(a.config, a.database)

	a.app.Use(a.logger)
	a.app.Use(routes.StaticPrefix, a.static)
	a.app.Static(routes.MediaPrefix, a.media.path, a.media.handler)

	a.middlewares = middlewares.New(a.app, a.session, a.config)
	a.router = routes.New(a.app)

	return a
}

func Init() *App {
	return &App{
		app: fiber.New(
			fiber.Config{
				AppName:      "goaccountant",
				ErrorHandler: handlers.DefaultErrorHandler,
				Views:        NewViews(),
				ViewsLayout:  "base",
			},
		),
		logger: NewLogger(),
		static: NewStatic(),
		media:  NewMedia(),
		config: config.New(),
	}
}

func (a *App) Listen() error {
	if a.config.UseTLS {
		return a.listenTLS()
	}
	return a.listenNoTLS()
}

func (a *App) listenTLS() error {
	go a.redirectServer()
	return a.app.ListenTLS(a.config.ConnectHTTPS, a.config.TLSCrt, a.config.TLSKey)
}

func (a *App) listenNoTLS() error {
	return a.app.Listen(a.config.ConnectHTTP)
}

func (a *App) redirectServer() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect(a.config.HTTPSRedirect)
	})
	if err := app.Listen(a.config.ConnectHTTP); err != nil {
		log.Fatalln(err)
	}
}
