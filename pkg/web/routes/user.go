package routes

import (
	"github.com/Devil666face/goaccountant/pkg/web/handlers"
	"github.com/Devil666face/goaccountant/pkg/web/middlewares"
)

func (r *AppRouter) setUser() {
	user := r.router.Group("/user")
	user.Use(r.wrapper(middlewares.Auth))
	user.Use(r.wrapper(middlewares.Admin))

	user.Get(
		"/list",
		r.wrapper(handlers.UserListPage),
	).Name("user_list")

	user.Get(
		"/create",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserCreateForm),
	).Name("user_create")
	user.Post(
		"/create",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserCreate),
	)

	user.Get(
		"/:id<int>/edit",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserEditForm),
	).Name("user_edit")
	user.Post(
		"/:id<int>/edit",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserEdit),
	)
}
