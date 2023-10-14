package routes

import (
	"github.com/Devil666face/goaccountant/pkg/store/session"
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	// "github.com/Devil666face/goaccountant/pkg/web/middlewares"

	"github.com/gofiber/fiber/v2"
)

var (
	StaticPrefix = "/static"
	MediaPrefix  = "/media"
)

type AppRouter struct {
	router  fiber.Router
	session *session.Store
}

func New(router fiber.Router, session *session.Store) *AppRouter {
	r := AppRouter{
		router:  router,
		session: session,
	}
	r.SetAuth()
	r.SetUser()
	// r.SetFree()
	// r.WithAuth()
	return &r
}

func (r *AppRouter) SetAuth() {
	auth := r.router.Group("/auth")
	auth.Get("/login", handlers.Login).Name("login")
}

func (r *AppRouter) SetUser() {
	user := r.router.Group("/user")
	user.Get("/list", handlers.UserList).Name("user_list")
	user.Get("/create", handlers.UserCreateForm).Name("user_create")
	user.Post("/create", handlers.UserCreate)
}

// func (r *AppRouter) SetFree() {
// 	r.router.Get("/login", handlers.Login).Name("login")
// }

// func (r *AppRouter) WithAuth() {
// 	r.router.Use(middlewares.AuthMiddleware(r.session))
// 	r.router.Get("/", handlers.Index)
// }
