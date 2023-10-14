package routes

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/database"
	"github.com/Devil666face/goaccountant/pkg/store/session"
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
	middlewares []func(*fiber.Ctx, *config.Config, *database.Database, *session.Store) error
}

func New(router fiber.Router, cfg *config.Config, db *database.Database, s *session.Store) *AppRouter {
	r := AppRouter{
		router:   router,
		config:   cfg,
		database: db,
		session:  s,
		middlewares: []func(*fiber.Ctx, *config.Config, *database.Database, *session.Store) error{
			middlewares.AllowedHostMiddleware,
			middlewares.CsrfMiddleware,
			middlewares.HtmxMiddleware,
		},
	}
	r.SetMiddlewares()
	r.SetAuth()
	r.SetUser()
	return &r
}

func (r *AppRouter) Handler(f func(*fiber.Ctx, *config.Config, *database.Database, *session.Store) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return f(c, r.config, r.database, r.session)
	}
}

func (r *AppRouter) SetMiddlewares() {
	for _, middleware := range r.middlewares {
		r.router.Use(r.Handler(middleware))
	}
}

func (r *AppRouter) SetAuth() {
	auth := r.router.Group("/auth")
	auth.Get("/login", r.Handler(handlers.Login)).Name("login")
}

func (r *AppRouter) SetUser() {
	user := r.router.Group("/user")
	user.Get("/list", handlers.UserList).Name("user_list")
	user.Get("/create", handlers.UserCreateForm).Name("user_create")
	user.Post("/create", handlers.UserCreate)
}
