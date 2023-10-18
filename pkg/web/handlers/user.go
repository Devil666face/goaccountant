package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

func UserList(uof *web.Uof) error {
	if uof.Ctx().IsHtmx() {
		return uof.Ctx().RenderWithCtx("user_content", fiber.Map{})
	}
	return uof.Ctx().RenderWithCtx("user_list", fiber.Map{
		"Title": "List of users",
	}, "base")
}

func UserCreateForm(uof *web.Uof) error {
	return uof.Ctx().RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(uof *web.Uof) error {
	u := models.User{}
	if err := uof.Ctx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(); err != nil {
		return uof.Ctx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return nil
}
