package routes

import (
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/gofiber/fiber/v2"
)

var (
	StaticPrefix = "/static"
	MediaPrefix  = "/media"
)

type AppRouter struct {
	router fiber.Router
}

func New(router fiber.Router) *AppRouter {
	r := AppRouter{
		router: router,
	}
	r.SetFree()
	return &r
}

func (r *AppRouter) SetFree() {
	r.router.Get("/", handlers.Index)
}
