package middlewares

import (
	"github.com/Devil666face/goaccountant/pkg/config"
	"github.com/Devil666face/goaccountant/pkg/store/session"

	"github.com/gofiber/fiber/v2"
)

type Middlewares struct {
	router      fiber.Router
	config      *config.Config
	session     *session.SessionStore
	middlewares []func(*fiber.Ctx) error
}

func New(router fiber.Router, session *session.SessionStore, config *config.Config) *Middlewares {
	m := Middlewares{
		router:  router,
		config:  config,
		session: session,
	}
	m.middlewares = m.getMiddlewares()
	m.setMiddlewares()
	return &m
}

func (m *Middlewares) setMiddlewares() {
	for _, middleware := range m.middlewares {
		m.router.Use(middleware)
	}
}

func (m *Middlewares) getMiddlewares() []func(*fiber.Ctx) error {
	return []func(*fiber.Ctx) error{
		AllowedHostMiddleware(m.config),
		CsrfMiddleware(m.session),
	}

}
