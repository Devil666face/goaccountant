package web

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/Devil666face/goaccountant/pkg/web/routes"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	app         *fiber.App
	static      func(*fiber.Ctx) error
	media       *Media
	config      *config.Config
	database    *database.Database
	router      *routes.Router
	session     *session.Store
	tables      []any
	middlewares []func(*fiber.Ctx) error
}

func New() *App {
	a := &App{
		app: fiber.New(
			fiber.Config{
				AppName:      "goaccountant",
				ErrorHandler: handlers.DefaultErrorHandler,
				Views:        NewViews(),
				// ViewsLayout:  "base",
			},
		),
		static: NewStatic(),
		media:  NewMedia(),
		config: config.New(),
		tables: []any{
			&models.User{},
		},
		middlewares: []func(*fiber.Ctx) error{
			NewLogger(),
			NewRecover(),
			NewHelmet(),
			NewCompress(),
			NewEncryptCookie(),
		},
	}
	a.setDefaultMiddlewares()
	a.setStores()
	a.setStatic()
	a.setRoutes()
	return a
}

func (a *App) setStores() {
	a.database = database.New(a.config, a.tables)
	a.session = session.New(a.config, a.database)
}

func (a *App) setStatic() {
	a.app.Use(routes.StaticPrefix, a.static)
	a.app.Static(routes.MediaPrefix, a.media.path, a.media.handler)
}

func (a *App) setRoutes() {
	a.router = routes.New(a.app, a.config, a.database, a.session)
}

func (a *App) setDefaultMiddlewares() {
	for _, m := range a.middlewares {
		a.app.Use(m)
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
		slog.Error(fmt.Sprintf("Start redirect server: %s", err))
		//nolint:revive //If connection for redirect server already busy - close app
		os.Exit(1)
	}
}
