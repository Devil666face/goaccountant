package routes

import (
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"
)

func (r *Router) setIndex() {
	index := r.router.Group("/")
	index.Use(r.wrapper(middlewares.Auth))

	index.Get(
		"",
		r.wrapper(handlers.Index),
	).Name("index")
}
