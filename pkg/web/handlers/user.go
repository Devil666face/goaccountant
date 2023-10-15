package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/Devil666face/goaccountant/pkg/web/view"

	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	if vc := view.New(c); vc.IsHtmx() {
		return vc.RenderWithCtx("user_content", fiber.Map{})
	}
	return view.New(c).RenderWithCtx("user_list", fiber.Map{
		"Title": "List of users",
	}, "base")
}

func UserCreateForm(c *fiber.Ctx) error {
	return view.New(c).RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(c *fiber.Ctx) error {
	u := models.User{}
	if err := c.BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(); err != nil {
		return view.New(c).RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return nil
}
