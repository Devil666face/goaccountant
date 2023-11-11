package routes

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"

	"github.com/gofiber/fiber/v2"
)

var (
	StaticPrefix = "/static"
	MediaPrefix  = "/media"
)

type Router struct {
	router      fiber.Router
	config      *config.Config
	database    *database.Database
	session     *session.Store
	middlewares []func(*web.Unit) error
}

func New(_router fiber.Router, _config *config.Config, _database *database.Database, _session *session.Store) *Router {
	r := Router{
		router:   _router,
		config:   _config,
		database: _database,
		session:  _session,
		middlewares: []func(*web.Unit) error{
			middlewares.Logger,
			middlewares.Recover,
			middlewares.Compress,
			middlewares.Limiter,
			middlewares.AllowHost,
			middlewares.SecureHeaders,
			middlewares.EncryptCookie,
			middlewares.Csrf,
			middlewares.Htmx,
		},
	}
	r.setMiddlewares()
	r.setAuth()
	r.setUser()
	r.setIndex()
	return &r
}

func (r *Router) wrapper(handler func(*web.Unit) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return handler(web.NewUnit(c, r.database, r.config, r.session))
	}
}

func (r *Router) setMiddlewares() {
	for _, middleware := range r.middlewares {
		r.router.Use(r.wrapper(middleware))
	}
}
