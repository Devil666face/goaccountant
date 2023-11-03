package routes

import (
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
)

func (r *AppRouter) setAuth() {
	auth := r.router.Group("/auth")

	auth.Get(
		"/login",
		r.wrapper(handlers.LoginPage),
	).Name("login")
	auth.Post(
		"/login",
		r.wrapper(handlers.Login),
	)
}
