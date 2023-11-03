package handlers

import (
	"fmt"

	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

func UserListPage(uof *web.Uof) error {
	if uof.ViewCtx().IsHtmx() && uof.ViewCtx().IsHtmxCurrentURL() {
		return uof.ViewCtx().RenderWithCtx("user_content", fiber.Map{
			"Users": models.GetAllUsers(uof.Database()),
		})
	}
	return uof.ViewCtx().RenderWithCtx("user_list", fiber.Map{
		"Title": "List of users",
		"Users": models.GetAllUsers(uof.Database()),
	}, "base")
}

func UserCreateForm(uof *web.Uof) error {
	return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(uof *web.Uof) error {
	u := models.User{}
	if err := uof.ViewCtx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	if err := u.Create(uof.Database()); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
		"Success": fmt.Sprintf("User %s - created", u.Username),
	})
}
