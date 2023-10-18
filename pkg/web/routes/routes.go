package routes

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"

	"github.com/gofiber/fiber/v2"
)

var (
	StaticPrefix = "/static"
	MediaPrefix  = "/media"
)

type AppRouter struct {
	router      fiber.Router
	config      *config.Config
	database    *database.Database
	session     *session.Store
	middlewares []func(*web.Uof) error
}

func New(router fiber.Router, cfg *config.Config, db *database.Database, s *session.Store) *AppRouter {
	r := AppRouter{
		router:   router,
		config:   cfg,
		database: db,
		session:  s,
		middlewares: []func(*web.Uof) error{
			middlewares.AllowedHostMiddleware,
			middlewares.CsrfMiddleware,
			middlewares.HtmxMiddleware,
		},
	}
	r.setMiddlewares()
	r.setAuth()
	r.setUser()
	return &r
}

func (r *AppRouter) handler(f func(*web.Uof) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		uof := web.NewUof(c, r.database, r.config, r.session)
		return f(uof)
	}
}

func (r *AppRouter) setMiddlewares() {
	for _, middleware := range r.middlewares {
		r.router.Use(r.handler(middleware))
	}
}

func (r *AppRouter) setAuth() {
	auth := r.router.Group("/auth")
	auth.Get("/login", r.handler(handlers.Login)).Name("login")
}

func (r *AppRouter) setUser() {
	user := r.router.Group("/user")
	user.Get("/list", r.handler(handlers.UserList)).Name("user_list")
	user.Get("/create", r.handler(handlers.UserCreateForm)).Name("user_create")
	user.Post("/create", r.handler(handlers.UserCreate))
}
